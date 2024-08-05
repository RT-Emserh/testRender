package send

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// bom aqui vamos fazer uma função auxiliar para verificar o valor de retorno de cada chamada amqp
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Send(b string) {
	// Bom aqui estamos nos conectando a função principal
	conn, err := amqp.Dial("amqp://joao:evilasio22@testrender-1-biqp.onrender.com:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	/*
		A conexão abstrai a conexão do socket e cuida da negociação e autenticação da versão do protocolo e assim por diante
		para nós. Em seguida, criamos um canal, que é onde a maior parte da API para fazer as coisas reside:
	*/
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// declaração da fila
	// bom o ch.queuedeclare declara uma nova fila do rabbitmq
	q, err := ch.QueueDeclare(
		// nome
		"hello", // aqui falamos o nome da fila que está sendo criada.
		// durable
		false, // A fila não é durável, o que significa que ela não sobreviverá a uma reinicialização do servidor RabbitMQ.
		// delete when unused
		false, //  A fila será deletada quando não estiver mais em uso.
		// exclusive
		false, // A fila não é exclusiva para a conexão atual. Pode ser acessada por outras conexões.
		// no-wait
		false, // Não esperar para que a declaração da fila seja confirmada.
		// arguments
		nil, // Nenhum argumento adicional está sendo passado para a declaração da fila.
	)

	// bom aqui so tratamos o erro
	failOnError(err, "Failed to declare a queue")
	// Criação do Contexto com Timeout
	// bom aqui criamos um novo contexto com um timeout de 5 segundos se ultrapassar esse tempo
	// a operação é cancelada para dar lugar a outra operação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// aqui com o defer cancel garante que o cancelamento do contexto será chamado quando a função terminar
	// liberando os recursos associados ao contexto
	defer cancel()
	// bom aqui preparamos para o envio da mensagem
	// aqui vamos passar um body como uma string
	body := b

	// Bom esse  PublishWithContext envia uma mensagem para a fila
	err = ch.PublishWithContext(ctx,
		// O nome da exchange. Aqui está vazio, o que significa que a mensagem será enviada diretamente para a fila.
		"", // exchange
		// A chave de roteamento, que é o nome da fila para onde a mensagem será enviada.
		q.Name, // routing key
		// false: mandatory é false, indicando que se a mensagem não puder ser roteada, ela não deve ser retornada.
		false, // mandatory
		// false: immediate é false, indicando que a mensagem não deve ser entregue imediatamente, se possível.
		false, // immediate
		// amqp.Publishing: Estrutura que contém as informações da mensagem.
		amqp.Publishing{
			// aqui no contentType falamos o tipo de conteudo, estamos determinando que é um texto simples
			ContentType: "text/plain",
			// bom aqui convertemos o corpo da mensagem para um slice de bytes
			Body: []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	// log.Printf(" [x] Sent %s\n", body)
}
