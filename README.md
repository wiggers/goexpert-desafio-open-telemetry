Pequeno sistema que recebe um CEP e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin) juntamente com a cidade. Para isto foi implementando um cenário em que temos o serviço 1, responsavel pelo input e temos o serviço 2 responsavel pela consulta do clima a partir do cep. 

Foi utilizado openTelemetry para retirar as métricas e tracing das aplicações.  

**Como Executar Localmente**
	
	Porta MS-1: 8000
	Porta MS-2: 8080

Executar o comando: 
	
		docker-compose up -d

Executar ambos os projetos:<br/> 
	pasta cmd no projeto MS-1 (ms-1/cmd) : <code>go run main.go</code> <br/> 
	pasta cmd no projeto MS-2 (ms-2/cmd) : <code>go run main.go</code>


Executar a chamada para o endpoint no **MS-1** (exemplo existente na pasta API do MS-1): 
	
	(POST) - localhost:8000/temperature
	
    ex: POST http://localhost:8000/temperature
            {
                "cep":"01153000"
            }

 
Acessar o **Zipkin** ou **Jaeger** para verificar o tracing através do endereço: 
	
	http://localhost:9411/zipkin/   
 	http://localhost:16686/search/
	 

Já as métricas estão sendo colhidas pelo **Prometheus** e podem ser analisadas através do **Grafana** bastando apenas entrar no Grafana através do link abaixo e adicinando o datasource.

	http://localhost:3000/

	

    




