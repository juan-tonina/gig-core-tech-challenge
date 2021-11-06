import * as Amqp from "amqp-ts";

const connection = new Amqp.Connection("amqp://guest:guest@host.docker.internal:5672/");
const exchange = connection.declareExchange("gig-exchange");
const queue = connection.declareQueue("gig-queue", {durable: false});
queue.bind(exchange);

export function pushToQueue(message: string) {
    exchange.send(new Amqp.Message(message));
}
