#### 18/09/2023

CURSO Go: Orientação a Objetos

@01-Minha primeira struct

@@01
Introdução

Olá, meu nome é Guilherme Lima e serei o instrutor neste curso de Go.
No curso aprenderemos a trabalhar com os principais tipos da linguagem e ainda com structs, ponteiros, referências e funções.

Veremos funções com um retorno só e outras com múltiplos retornos e saberemos como utilizá-las e trabalhar com métodos.

Também aprenderemos a trabalhar com composição, encapsulamento e interface.

O foco principal do curso será a linguagem Go.

Então, se quiser aperfeiçoar o conhecimento nessa linguagem incrível, convido a participar das aulas.

Ainda não veremos o tema concorrência durante o treinamento, pois nosso alvo são as pessoas que ainda não têm muita familiaridade com o Go e querem se aprofundar.

@@02
Variáveis e tipos

Um certo banco está com dificuldade em gerenciar suas contas correntes. Vamos auxiliá-lo criando uma aplicação em Go para a gestão das contas.
Para começar, acessaremos o nome de usuário de nosso computador e a pasta "go", em que foi instalada nossa aplicação Go. Temos três pastas e iremos em "src > github.com" e em seguida o nome do usuário que foi definido. Haverá todos os projetos criados utilizando a linguagem e criaremos um novo clicando com o botão direito e selecionando "New folder". Chamaremos esse projeto de "banco".

Usaremos o Visual Studio Code para editar o código. A utilização deste editor é aconselhável para acompanhar as aulas, apesar de qualquer outro poder ser usado também.

No VS Code clicaremos na opção "Open folder..." para abrir uma pasta. Selecionaremos "banco" e fecharemos a página de boas-vindas do programa. Ainda não há nada em nosso arquivo. Clicaremos na opção "New file", no ícone da folha em branco, e nomearemos como main.go. Todo código Go precisa de um pontapé inicial, um código que seja o ponto de partida de nossa aplicação.

Só de termos nomeado o arquivo, o compilador já sublinhará a tela de vermelho porque precisamos dizer qual é o pacote do nosso código. Então, dentro do nosso folder de "banco" teremos o arquivo main.go, e nele, precisaremos de uma função inicial. Chamaremos a função de main. Então nossa aplicação começa a partir desse código.

Dentro das chaves que abriremos depois da função, vamos inserir o código que será executado na função principal.

package main

func main(){

}COPIAR CÓDIGO
Nossa aplicação, como já foi dito, servirá para gerenciar contas correntes. Antes, precisamos saber quais campos temos nessas contas correntes. Pensando nisso, o banco já nos mandou um arquivo nos mostrando quais são os dados que guardam para manter as informações nessa conta corrente.

Uma conta corrente nesse banco tem o nome do Titular responsável pela conta, o número da Agência, o Número da Conta e o Saldo. Criaremos uma forma de guardar esses valores na memória do computador. Para guardar valores na memória, devemos criar uma variável. Começaremos com a criação da variável para o Titular da conta.

Dentro da função, escreveremos var e o nome da variável titular. Queremos armazenar nela uma palavra, referente ao nome de uma pessoa. Nesse caso, escrevemos uma palavra reservada no código chamada string. Armazenaremos nessa string o código "Guilherme". Então, criamos uma variável chamada titular do tipo string e atribuímos a ela utilizando um sinal de igual (=) a palavra "Guilherme". Colocamos "Guilherme" entre aspas porque as utilizamos no Go para a linguagem saber que dentro das aspas há uma palavra.

Tendo criado o Titular, criaremos também a Agência com var numeroAgencia. Porém, o Número da Agência não será uma palavra. Serão os números "589" que armazenaremos na variável e precisamos identificar o tipo. Para números, utilizaremos o tipo int. e não será necessário o uso de aspas no código. var numeroAgencia int = 589.

Agora, vamos gerar o Número da Conta com var numeroConta. É importante observar que estamos digitando a primeira letra minúscula, mas escrevemos as primeiras letras das próximas palavras em maiúsculas. O nome técnico desta convenção é camelCase. Como numeroConta armazenará números também, usaremos igualmente o tipo int, e teremos var numeroConta int = 123456.

O saldo, porém, será um número diferente. Não queremos que essa variável receba um número inteiro e sim um decimal fracionário, para que ele possa ter centavos. Para isso utilizamos a palavra float. O autocomplete do VS Code mostrará que temos dois tipos de float, float32 e float64. Ainda não vamos nos preocupar em diferenciá-los, mas escolheremos o tipo float64. Para testar, colocaremos um número fracionário na conta, var saldo float64 = 125,50. No entanto, não costumamos usar vírgulas, e sim o ponto.

package main

func main(){
    var titular string = "Guilherme"
    var numeroAgencia int = 589
    var numeroConta int = 123456
    var saldo float64 = 125.50

}COPIAR CÓDIGO
Salvaremos o código e veremos que vários erros terão sido sublinhados de vermelho. Os alertas serão que as variáveis foram declaradas e não foram usadas. Isso porque em Go não podemos criar uma variável e não utilizá-la para nada.

Vamos exibir esses campos na tela, ou seja, o que há nas variáveis titular, numeroConta e saldo. Digitaremos fmt, um outro pacote capaz de escrever na tela para nós, e usaremos uma propriedade .Println(). Nos parênteses passaremos as variáveis.

fmt.Println(titular, numeroAgencia, numeroConta, saldo)

Assim que salvamos, o sublinhado vermelho ficará verde e logo após o package main aparece um import "fmt", pois a biblioteca foi importada para nós. A primeira letra de Println() precisa sempre ser maiúscula.

Sendo assim, foi feita a impressão. Ainda é necesário visualizá-la. Clicaremos com o botão direito em "main.go" e selecionaremos a opção "Open in Terminal", ou abrir no terminal. Será aberto um pequeno terminal na parte inferior da tela.

Se digitarmos nesse terminal o comando ls e pressionarmos "Enter", será mostrado no terminal que temos um arquivo main.go conforme o esperado. Solicitamos a execução desse arquivo por meio de um go run e do nome do arquivo que queremos executar, ou seja, go run main.go. Agora leremos todas as informações que estão armazenadas dentro das variáveis. "Guilherme 589 123456 125.5", omitindo o "0" do valor do saldo.

Essas são as informações referentes as variáveis titular, numeroAgencia, numeroConta, saldo.

Criaremos um outro titularvar titular string "Luciene". Antes de criar os outros campos, vamos visualizar no console com fmt.Println(titular) e salvar. No entanto, quando salvarmos, aparecerá uma mensagem de erro. Ela nos avisará que titular já foi declarada no bloco anterior. Sendo assim, não podemos declarar outra variável de mesmo nome e tentarmos atribuir um valor diferente.

Por isso chamaremos a nova variável de titular2

var titular2 string = "Luciene"

fmt.Println(titular2)COPIAR CÓDIGO
Limparemos o terminal pressionando "Ctrl + L" e digitaremos go run main.go. Será exibido o nome "Luciene" além de todos os valores das variáveis declaradas anteriormente,

Então sempre precisaremos criar outra variável com outro nome para adicionar outra informação. Faremos o mesmo para as demais

var numeroAgencia2 int = 555
var numeroConta2 int= 111333
var saldo2 float64 = 200

fmt.Println(titular2, numeroagencia2, numeroConta2, saldo2)COPIAR CÓDIGO
Vamos executar e rodar e tudo dará certo, mas a manutenção do código está se tornando cada vez mais difícil. Temos dois titulares, Guilherme e Luciene. A medida que o código cresce, teremos mais variáveis e a organização se tornará mais complexa. Pensando nisso, no próximo vídeo veremos uma forma de organizar melhor as contas correntes.

@@03
Structs

Vimos que uma conta corrente possui um Titular, uma Agência, um Número da Conta e um Saldo.
Quando criarmos outras contas correntes, elas terão esses mesmos campos e precisaremos declarar diversas variáveis. Podemos evitar esse trabalho, já que em Go temos uma forma de estruturar esses campos.

Criaremos uma estrutura de conta corrente com todos os campos necessários e quando criarmos a conta tanto do Guilherme quanto da Luciene conseguiremos utilizar a mesma estrutura com os mesmos campos.

Para criar a estrutura no Go vamos ao código e removeremos a segunda conta criada, deixando apenas a primeira para fins de comparação. Pressionaremos "Command + J" para remover nosso terminal.

Queremos criar uma estrutura chamada ContaCorrente e ela terá as variáveis titular, numeroAgencia, numeroConta e saldo. Mas se só declararmos as variáveis, o código não entenderá que essa é uma estrutura. Precisamos identificá-la e especificar os tipos de variáveis. uma string, um int, outro int e um float64.

Para dizermos para a linguagem de fato que ContaCorrente é uma estrutura, utilizamos um prefixo antes do nome dela, e depois um sufixo. O prefixo indicará que isso é um novo tipo, por isso usaremos a palavra type. Após ContaCorrente escreveremos struct, ou seja, "estrutura" em Inglês.

Devemos identificar que todos esses campos se tratam da struct, por isso vamos envolvê-los entre chaves.

type ContaCorrente struct {
    var titular       string
     var numeroAgencia int
    var numeroConta   int
    var saldo         float64
}COPIAR CÓDIGO
Pressionaremos "Command + S" para salvar, mas veremos uma marcação de erro em todos os nossos campos. O alerta será de que a palavra var não será necessária nesse contexto, pois não precisamos identificar que os campos são variáveis. Removeremos var de todos os campos antes de salvar novamente.

type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}COPIAR CÓDIGO
Agora não teremos mais nenhuma indicação de erro. Criamos nossa primeira struct, um tipo que conterá vários outros tipos dentro dele. As contas dos clientes, tanto Guilherme quanto Luciene, deverão utilizar a estrutura.

Não precisaremos mais exibir todos os campos no print. Vamos passar para fmt.Println() que a partir de agora queremos trabalhar com ContaCorrente. Para deixar claro que se trata de uma estrutura, vamos abrir e fechar chaves na sequência. Se passarmos o mouse sobre ContaCorrente{} será mostrada a estrutura que criamos desse tipo.

func main() {
    fmt.Println(ContaCorrente{})
}COPIAR CÓDIGO
Vamos salvar e teclaremos "Command + J" para abrir o terminal mais uma vez. Limparemos o terminal com "Ctrl + L" e digitaremos go run main.go para rodar. Ele exibirá uns números estranhos: { 0 0 0}. Não passamos nenhuma informação e mesmo assim foram impressos esses valores.

Apesar de não termos inserido nenhum valor, os elementos ganham automaticamente um valor inicial. Esse valor é chamado de zero value ou inicialização zero. Então, pro que for string, será adicionado um vazio. Para o que for um inteiro ou do tipo float, será atribuído um "0".

Porém, não é o que queríamos. A primeira conta corrente não deve ser uma string vazia e três zeros. Ela deve ter exatamente as informações referentes ao Guilherme.

Faremos essas atribuições no próximo vídeo.

@@04
Utilizando a struct

Quando não passamos nenhuma informação para a struct, ela é criada automaticamente com os valores padrão. Mas não é o que queremos, nossa ideia é criar uma conta corrente com todos os campos correspondentes.
Primeiro vamos lembrar que a primeira palavra que usamos para declarar nossa struct é type, assim como os tipos que declaramos para as variáveis serem strings ou inteiros.

Agora criaremos uma nova variável na func main(). Para isso precisaremos alocar um espaço na memória por meio de var e chamaremos a variável de contaDoGuilherme. Ela não será de um tipo inteiro, string ou float, mas do tipo contaCorrente. Por esse motivo utilizamos a palavra type na declaração de contaCorrente, uma estrutura que tem vários campos.

Vamos exibir contaDoGuilherme no terminal na sequência.

func main() {
    var contaDoGuilherme ContaCorrente = ContaCorrente{}

    fmt.Println(contaDoGuilherme)
}COPIAR CÓDIGO
Se eu pedir para exibir a conta no terminal, após salvar o código, limpar o terminal e executar mais uma vez, veremos o mesmo resultado de antes, um zero value ou a inicialização automática com valores que o Go coloca.

Mas a forma de escrita não está tão adequada a que costuma-se utilizar na linguagem Go. Quando queremos alocar algo na memória e já fazer uma atribuição, podemos fazer uma atribuição curta. No lugar de usar o sinal de igual, colocaremos dois pontos antes do símbolo, apagaremos o tipo contaCorrente e poderemos apagar também a palavra var. Já sabemos que essa é uma nova variável e que estamos colocando um tipo dentro dela.

func main() {
    contaDoGuilherme := ContaCorrente{}

    fmt.Println(contaDoGuilherme)
}COPIAR CÓDIGO
Como não temos nenhum erro, rodaremos o código de novo e ainda vamos ver o mesmo resultado.

Mas não é nosso objetivo. Queremos obter todos os campos. Mas já temos a contaCorrente{}, abrimos e fechamos chaves. Podemos colocar dentro dessas chaves os campos e as informações do cliente.

func main() {
    contaDoGuilherme := ContaCorrente{titular: "Guilherme",
        numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}


    fmt.Println(contaDoGuilherme)
}COPIAR CÓDIGO
Vamos salvar e não haverá nenhuma mensagem de erro. Executaremos o código mais uma vez com go run main.go e será impresso "Guilherme 589 123456 125.5", as informações que queríamos.

Se quisermos criar uma outra conta, por exemplo, a conta da Bruna, precisamos declarar contaDaBruna como do tipo contaCorrente{} por meio dos sinais ":= " . Dentro das chaves já não precisaremos passar o nome do campo e o valor do campo. Já podemos passar os valores diretamente, na ordem em que foram criados. A titular será "Bruna", Agência 222, Conta 111222 e Saldo 200. Exibiremos o print da contaDaBruna na última linha do código.

func main() {
    contaDoGuilherme := ContaCorrente{titular: "Guilherme",
        numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

    contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

    fmt.Println(contaDoGuilherme)
    fmt.Println(contaDaBruna)
}COPIAR CÓDIGO
Salvaremos, vamos executar e teremos o mesmo resultado, as informações da Bruna de acordo com os campos.

Então, queremos usar sempre o segundo modo, pois com ele não precisaremos escrever o nome de todos os campos novamente. Algo interessante é que quando temos uma estrutura e vamos necessariamente todos os campos, podemos usar o segundo modo, mais sucinto, e será entendido que estamos passando as informações na ordem que passamos os campos a princípio.

Se não precisarmos de todos os campos, por exemplo, apenas titular e saldo, teríamos que escrevê-los para especificar que os valores se tratam dos dois. Na conta do Guilherme, vamos tirar numeroConta e numeroAgencia.

contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}

fmt.Println(contaDoGuilherme)COPIAR CÓDIGO
Vamos rodar e aparecerá no terminal o nome do Titular, Guilherme, no Número da Agência e da Conta terão sido colocados "0"s e o saldo aparecerá também.

Temos essas duas opções de utilização da nossa estrutura. Do outro modo, teríamos 8 variáveis referentes aos campos do Guilherme e da Bruna. Dessa forma, com poucas linhas de código conseguimos um resultado mais elegante usando a linguagem Go.

O código ficará disponível na última atividade da aula atual, em O que aprendemos? para a utilização.

@@05
Reflexão sobre structs
PRÓXIMA ATIVIDADE

Para criar uma nova struct, usamos o prefixo type, seguido do nome da struct, o sufixo struct e declaramos os campos entre chaves, conforme o exemplo abaixo:
type Pessoa struct {
    nome string
    idade int
}COPIAR CÓDIGO
Sabendo disso, analise as afirmações abaixo e marque as verdadeiras em relação ao uso de structs no Go.

As variáveis ficam mais organizadas.
 
Certo! As variáveis ficarão agrupadas para cada nova pessoa criada.
Alternativa correta
Permite reutilizar o código.
 
Certo! Não precisamos criar 2 variáveis para criar um nova pessoa.
Alternativa correta
Permite criar uma nova pessoa com o código Pessoa(nome: "Luísa", idade: 28)
 
Alternativa correta
Só permite criar um novo tipo informando o nome dos campos.

@@06
Para saber mais
PRÓXIMA ATIVIDADE

Inicialização zero e nil
Mesmo não provendo nenhum valor, o Go garante inicializar todas as variáveis, conforme a imagem abaixo:

Título da tabela: Zero-initialization. Primeira linha da tabela: false para bool para. Segunda linha: 0 para int. Terceira linha: 0 para float. Quarta linha: "" para string. Quinta linha: {} para struct

Porém, em muitas linguagens existe uma maneira de denotar um ponteiro nulo que, essencialmente, não aponta para nenhum lugar. Por exemplo: em C é NULL, em Python é None e em java é null. Em Go, temos o nil.

Nil e Inferência
Observe o seguinte exemplo:

package main

import (
    "fmt"
)

func main() {
    a := nil
    fmt.Println(a)
}COPIAR CÓDIGO
Será que vai compilar?

Não, não vai. O compilador imprimirá o seguinte erro: use of untyped nil, que significa uso não tipado do nil.

Aqui estamos tentando atribuir um valor nil apontando para algum lugar sem fornecer seu tipo e esperamos que o compilador deduza isso. O compilador não sabe se esta variável é um inteiro, uma string, um array ou uma structure.

Neste link, você pode executar o código acima.

Nil com um tipo definido
Sabendo disso, observe o exemplo abaixo onde apontamos para um tipo definido:

package main

import (
    "fmt"
)

func main() {
    var s *string = nil
    fmt.Println(s)
}COPIAR CÓDIGO
Neste caso, o programa compila e retorna <nil> como esperado.

Neste link, você pode executar o código acima.

Comparando zero value
Para finalizar, observe o seguinte programa, no qual criamos duas variáveis: uma float64 e uma int, sem atribuir valor, e as comparamos:

package main

import (
    "fmt"
)

func main() {
    var f float64
    var i int 

    fmt.Println(f==i)
}COPIAR CÓDIGO
Neste link, você pode executar o código acima.

Conclusão
Recebemos uma mensagem com um erro informando que os tipos são incompatíveis. Não podemos comparar o valor atribuído pela inicialização zero se temos tipos diferentes. Portanto, por mais que o Go garanta a inicialização zero de diferentes tipos, devemos ficar atentos com os tipos que estamos trabalhando.

https://play.golang.org/p/kVcneO6nVdS

https://play.golang.org/p/gkwKo7rholt

https://play.golang.org/p/Ri6dLuyhjeg

@@07
Faça como eu fiz na aula
PRÓXIMA ATIVIDADE

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito, excelente. Se ainda não, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código aqui escrito.
Se por acaso você já domina essa parte, em cada capítulo, você poderá baixar o projeto feito até aquele ponto.

O gabarito deste exercício é o passo a passo demonstrado no vídeo. Tenha certeza de que tudo está certo antes de continuar. Ficou com dúvida? Podemos te ajudar pelo nosso fórum.

@@08
O que aprendemos?
PRÓXIMA ATIVIDADE

Nessa aula:
Criamos nossa primeira struct chamada ContaCorrente;
Em seguida criamos um tipo com base na struct Conta corrente sem atribuir valor para os campos e descobrimos o Zero-initialization ou Inicialização zero;
Para finalizar, vimos duas formas utilizar a struct criada: informando ou não os nomes dos campos.
Projeto desenvolvido nesta aula
Neste link, você fará o download do projeto feito até esta aula.

Caso queira visualizar o código desenvolvido até aqui, clique neste link.

Na próxima aula
Vamos entender na prática o que são ponteiros, aprender uma outra forma de utilizar a struct e criar um método que possibilita o saque!

https://github.com/alura-cursos/go_oo/archive/master.zip

https://github.com/alura-cursos/go_oo

