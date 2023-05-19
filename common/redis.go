package common

// Consume 返回值代表消费过程中遇到的无法处理的错误
// func (q *ACKListMQ) Consume(ctx context.Context, topic string, h Handler) error {
// 	for {
// 		// 获取消息
// 		body, err := q.client.LIndex(ctx, topic, -1).Bytes()
// 		if err != nil && !errors.Is(err, redis.Nil) {
// 			return err
// 		}
// 		// 没有消息了，休眠一会
// 		if errors.Is(err, redis.Nil) {
// 			time.Sleep(time.Second)
// 			continue
// 		}
// 		// 处理消息
// 		err = h(&Msg{
// 			Topic: topic,
// 			Body:  body,
// 		})
// 		if err != nil {
// 			continue
// 		}
// 		// 如果处理成功，删除消息
// 		if err := q.client.RPop(ctx, topic).Err(); err != nil {
// 			return err
// 		}
// 	}
// }
