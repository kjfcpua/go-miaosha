/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/23 12:15 下午
 * @Description: RabbitMQ封装 https://www.rabbitmq.com/getstarted.html
 */
package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// amqp://账号:密码@rabbitmq服务器地址:端口号/virtual hosts
const MqUrl = "amqp://miaosha:gomiaosha@127.0.0.1:5672/miaosha"

// 定义RabbitMQ对象
type RabbitMQ struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	QueueName    string // 队列名称
	ExchangeName string // 交换机名称
	RoutingKey   string // key名称
	MqUrl        string // 连接信息
}

// 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.connection.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

// 创建RabbitMQ结构体实例
func NewRabbitMQ(queueName, exchangeName, routingKey string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{
		QueueName:    queueName,
		ExchangeName: exchangeName,
		RoutingKey:   routingKey,
		MqUrl:        MqUrl,
	}
	var err error
	rabbitMQ.connection, err = amqp.Dial(rabbitMQ.MqUrl)
	rabbitMQ.failOnErr(err, "创建连接失败")
	rabbitMQ.channel, err = rabbitMQ.connection.Channel()
	rabbitMQ.failOnErr(err, "获取Channel失败")
	return rabbitMQ
}

// 创建Simple简单模式RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

func (r *RabbitMQ) applyQueueArgs(exclusive bool) {
	// 申请队列，如果队列不存在则创建，如果存在则调过创建
	// 保证队列存在，消息能发送到队列中
	q, err := r.channel.QueueDeclare(
		r.QueueName, // name
		false,       // 消息是否持久化
		false,       // 最后一个监听失效是否自动删除消息
		exclusive,   // 是否具有排他性（其他用户是否可见）
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	r.QueueName = q.Name
	r.failOnErr(err, "申请队列失败")
}

// 试探性申请队列
func (r *RabbitMQ) ApplyQueue() *RabbitMQ {
	r.applyQueueArgs(false)
	return r
}

// 发送消息
func (r *RabbitMQ) publish(message string) {
	err := r.channel.Publish(
		r.ExchangeName, // 交换机
		r.RoutingKey,
		// mandatory若为true,根据Exchange类型和routKey规则，
		// 如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// immediate若为true,当exchange发送消息到队列后发现队列上没有绑定消费者，
		// 则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "发送消息失败")
	log.Println("消息发送成功！")
}

// 生产简单模式消息
func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列
	r.ApplyQueue()
	// 发送消息到队列中
	r.publish(message)
}

// 消费消息
func (r *RabbitMQ) Consume() {
	// 申请队列（保证队列存在，消息能发到队列中），若不存在队列，会自动创建，已存在则跳过创建
	msgs, err := r.channel.Consume(
		r.QueueName, // queue
		"",          // 用来区分多个消费者
		true,        // 是否自动应答，消费完成通知RabbitMQ删除该条消息
		false,       // 是否具有排他性（其他用户是否可见）
		false,       // 如果设置为true,表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,       // 队列消息是否阻塞
		nil,         // 其他参数
	)
	r.failOnErr(err, "注册消费者失败")

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("接收到消息 : %s", string(d.Body))
		}
	}()
	log.Printf("消费者正在等待接收消息")
	<-forever
}

// 创建订阅模式RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ("", exchangeName, "")
	return rabbitMQ
}

// 试探性创建交换机(广播类型:"fanout",路由模式："direct")
func (r *RabbitMQ) applyExchange(kind string) {
	err := r.channel.ExchangeDeclare(
		r.ExchangeName, // name
		kind,           // type 交换机类型（fanout：广播类型）
		true,           // durable 是否持久化
		false,          // auto-deleted 是否自动删除
		//true表示这个exchange不可以被client用来推送消息，
		// 仅用来进行exchange和exchange之间的绑定
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	r.failOnErr(err, "创建交换机失败")
}

// 队列绑定到交换机
func (r *RabbitMQ) bindingQueueExchange() {
	err := r.channel.QueueBind(
		r.QueueName,    // queue name
		r.RoutingKey,   // routing key,订阅模式下key必须为空
		r.ExchangeName, // exchange
		false,
		nil,
	)
	r.failOnErr(err, "队列绑定到交换机失败")
}

// 订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	// 试探性创建交换机
	r.applyExchange("fanout")
	r.publish(message)
}

func (r *RabbitMQ) ReceiverSub() {
	// 试探性创建交换机
	r.applyExchange("fanout")
	// queueName=""表示队列随机生成
	r.applyQueueArgs(true)
	r.bindingQueueExchange()
	// 消费消息
	r.Consume()
}

// 创建路由模式RabbitMQ实例
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ("", exchangeName, routingKey)
	return rabbitMQ
}

func (r *RabbitMQ) PublishRouting(message string) {
	r.applyExchange("direct")
	r.publish(message)
}

func (r *RabbitMQ) ReceiverRouting() {
	r.applyExchange("direct")
	r.applyQueueArgs(true)
	r.bindingQueueExchange()
	r.Consume()
}

// 创建话题模式RabbitMQ实例
func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	return rabbitmq
}

// 话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) {
	r.applyExchange("topic")
	r.publish(message)
}

// 话题模式接收消息
func (r *RabbitMQ) ReceiverTopic() {
	r.applyExchange("topic")
	r.applyQueueArgs(true)
	r.bindingQueueExchange()
	r.Consume()
}
