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

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito, excelente. Se ainda não, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código aqui escrito.
Se por acaso você já domina essa parte, em cada capítulo, você poderá baixar o projeto feito até aquele ponto.

O gabarito deste exercício é o passo a passo demonstrado no vídeo. Tenha certeza de que tudo está certo antes de continuar. Ficou com dúvida? Podemos te ajudar pelo nosso fórum.

@@08
O que aprendemos?

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

#### 19/09/2023

@02-Referência, ponteiro e métodos

@@01
New e ponteiros

Vimos duas formas de utilizar a struct, a primeira passando o campo e o valor que queremos armazenar dentro dessas variáveis e a segunda passando o conteúdo sem especificar os campos, desde que eles estejam na ordem em que foram declarados.
Será que só existem essas formas de utilizar a struct? Não, há a opção de usá-la de forma similar a outras linguagens de programação como Java e C#. Aprenderemos como.

Nosso primeiro passo será declarar uma variável chamada contaDaCris. Ela será do tipo ContaCorrente. Na linha seguinte, faremos contaDaCris = new(). A palavra new é bastante conhecida de quem já programa em Java ou C#. Dentro dos parênteses passaremos o tipo, que é ContaCorrente.

A partir de então, conseguiremos atribuir os valores acrescentando um ponto (.) após contaDaCris e o nome do campo na sequência. Então atribuiremos o valor desse campo.

var contaDaCris ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"

fmt.Println(contaDaCris)COPIAR CÓDIGO
Teclaremos "Command + J" para abrir o terminal e "Ctrl + L" para limpá-lo. O "Command + S" será necessário para salvar o código. Porém, aparecerá um erro, pois não podemos utilizar nosso tipo ContaCorrente já atribuindo um valor. Isso porque temos uma variável do tipo ContaCorrente, a contaDaCris. Mas o código não entendeu que o tipo da variável é o mesmo que está sendo passado para new(). Precisamos identificar que se tratam do mesmo tipo.

Para conseguir dizer que contaDaCris aponta para uma ContaCorrente, colocaremos um asterisco na frente. Teremos var contaDaCris *ContaCorrente.

Agora não teremos mais nenhum erro. Limparemos o terminal e vamos roda de novo. Veremos no terminal:

&{Cris 0 0 0}

Esse resultado se deve a só termos atribuído o nome, então os outros valores tomaram a forma de zero velho. Atribuiremos agora os demais valores.

Mas por que o "&" e por que temos que apontar para ContaCorrente com asterisco? Nas outras alternativas de uso da struct, o código entendia que que a conta corrente do cliente apontava para um tipo, uma estrutura de conta.

Exemplificaremos o que está acontecendo. Vamos imaginar que um edifício tem uma cobertura super bonita, apertamentos maiores no penúltimo e no último andar, e o térreo e o primeiro andar têm todos os apartamentos do mesmo tamanho.

Para cada apartamento há uma caixa de correio que os identifica. As correspondências que chegam para cada apartamento ficam nas respectivas caixas de correio. Eles possuírem tamanhos diferentes é um ponto importante.

Isso quer dizer que nossas caixas de correio são como os nossos ponteiros ou referências. Cada caixa, independentemente do tamanho do apartamento, apontará para o mesmo lugar se tratando do edifício. Os ponteiros também serão todos iguais. Porém, quando falamos de cada apartamento, os ponteiros apontarão para lugares diferentes do prédio.

Para concluir, as caixas do correio serão nossos ponteiros. O edifício será a memória da nossa aplicação, o local em que podemos armazenar informações. Cada apartamento corresponde a um tipo a ser armazenado.

Então existem tipos menores, os relacionados com os apartamentos do térreo e primeiro andar, onde não precisamos de muito espaço para o armazenamento. O segundo andar será semelhante.

No terceiro andar já teremos um apartamento maior, que permite guardar elementos maiores. O ponteiro, no entanto, será do mesmo tamanho. Essa é apenas uma reflexão simples do que ocorre no nosso desenvolvimento.

Então a contaDaCris precisa apontar para ContaCorrente. O código em Go não entenderá corretamente se tirarmos e asterisco e salvarmos a aplicação, pois ficará sem entender para onde new está apontando, se é a conta corrente da Cris ou uma nova conta. A partir do momento em que o ponteiro é colocado, é como se alocássemos um espaço e disséssemos que a caixa do correio aponta para um apartamento em particular.

A caixa de correio contaDaCris apontará para o apartamento ContaCorrente . Podemos criar cada um dos campos da conta corrente de Cris de acordo com a struct. Criaremos o saldo nesse momento, com contadaCris.saldo = 500.

O último detalhe será o "&" que aparece no terminal quando fazemos a impressão do que temos. Isso porque seguindo o mesmo raciocínio do nosso exemplo, &{Cris 0 0 500} indica que os campos são um conteúdo que está dentro do apartamento. Mas "&" não é interessante para nós, somente o conteúdo.

Por isso, colocaremos o asterisco em contaDaCris também na hora da impressão.

var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.saldo = 500


fmt.Println(*contaDaCris)COPIAR CÓDIGO
Se limparmos o terminal e rodarmos de novo, o resultado será igual ao que teríamos usando as outras opções para structs, {Cris 0 0 500} somente.

Então, por debaixo dos panos ocorrerá todo o processo referente aos ponteiros e à alocação da memória nos outros códigos também.

A pergunta que fica é: com qual forma de usar structs é mais fácil trabalhar. Podemos usar a técnica de escrever contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5} quando provavelmente não vamos precisar sempre de todos os campos preenchidos.

A segunda forma, contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200} é mais indicada quando todos os campos serão utilizáveis. Esses serão os modos mais compatíveis com o Go.

Mas o terceiro modo também é importante sabermos usar, pois se precisar lidar com um projeto com esse modelo, saberá do que se tratam os ponteiros e como manipular esses dados.

Essa foi uma aula para termos ideia de que existe essa alternativa e como utilizá-la, mas nos próximos momentos trabalharemos sempre com o primeiro ou o segundo jeitos de lidar com structs.

@@02
Comparando tipos

Vimos algumas formas de utilizar nossa struct. Mas para entender um pouco mais a fundo o conceito de ponteiro, faremos um teste juntos.
Selecionaremos todo o código a partir de contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200} até embaixo, vamos segurar a tecla "Command" e pressionar a barra "//". Assim comentaremos esse trecho do código e ele não será mais executado.

Tiraremos apenas o comentário de fmt.Println(contaDoGuilherme), pois queremos exibi-lo. Se salvarmos, não teremos nenhuma mensagem de erro, pois a criação de contaDoGuilherme já tinha sido feita anteriormente.

Agora criaremos outra conta. Vamos copiar a variável contaDoGuilherme com todas as informações de campos e chamar a duplicata de contaDoGuilherme2

contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

contaDoGuilherme2 := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}COPIAR CÓDIGO
Vamos salvar o código e veremos uma mensagem de erro nos dizendo que contaDoGuilherme2 foi usada mas não foi declarada. Podemos resolver isso acrescentando a variável à exibição em fmt.Println(contaDoGuilherme, contaDoGuilherme2).

Salvaremos e rodaremos no terminal com go run main.go. Será impresso o conteúdo de ambas as contas, que é o mesmo. Mas não queremos exibir o conteúdo. Vamos comparar se elas de fato são iguais, e para isso usaremos o sinal de "==".

fmt.Println(contaDoGuilherme == contaDoGuilherme2)

Limparemos o terminal mais uma vez e quando executarmos, teremos como resposta a palavra "true", que significa "verdadeiro" em Inglês. Então, por mais que eu tenha um apartamento com um conteúdo e tenha outro apartamento com o mesmo conteúdo, o Go é inteligente o bastante para entender que não estamos tentando saber se as informações se referem a uma conta (apartamento) ou outra. Ele compara o conteúdo de ambos e avisa que o resultado da comparação é verdadeiro.

Mas vamos mudar um valor. Colocaremos "580" no numeroAgencia de contaDoGuilherme2, deixando "589" em contaDoGuilherme. Rodaremos a comparação novamente e dessa vez aparecerá a mensagem "false" no terminal, ou seja, falso, pois com a alteração não é mais verdadeiro que o conteúdo das duas variáveis é comparável.

Então, como fizemos esse procedimento de comparação para contaDoGuilherme, faremos o mesmo para contaDaBruna. Comentaremos os trechos do código referentes a contaDoGuilherme, apagaremos contaDoGuilherme2 e descomentar os demais.

Então, da mesma forma que foi feito anteriormente, no método que passa todos os campos e as informações referentes a eles, criaremos contaDaBruna2 com o mesmo conteúdo da primeira variável e vamos compará-las no print.

contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

fmt.Println(contaDaBruna == contaDaBruna2)
}
COPIAR CÓDIGO
Vamos executar o código e o terminal exibirá "true", já que o conteúdo de ambas as variáveis é igual. Se trocarmos o Número da Conta de contaDaBruna para "111" no lugar de "222", mantendo o de contaDaBruna2, e rodarmos novamente, aparecerá a palavra "false" no terminal, porque os conteúdos dos endereços agora são diferentes.

Então, o modo Go de escrever o código utilizando o sinal de dois pontos e igual (:=) e a ContaCorrente{} (com conteúdo entre chaves) entende que quando fazemos uma comparação queremos levar em comparação o conteúdo dos apartamentos.

Substituiremos a impressão por contaDoGuilherme e depois comentaremos.

// fmr.Print(contaDoGuilherme)

Agora faremos o mesmo para contaDaCris. Primeiramente criaremos contaDaCris2. Por fim, compararemos ambas as contas e imprimiremos o resultado.

var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.titular = 500

var contaDaCris2 *ContaCorrente
contaDaCris2 = new(ContaCorrente)
contaDaCris2.titular = "Cris"
contaDaCris2.titular = 500

fmt.Println(contaDaCris == contaDaCris2)COPIAR CÓDIGO
Salvaremos o código e vamos executar. No entanto, aparecerá no terminal que isso é falso, diferentemente dos outros dois casos em que havendo o mesmo conteúdo nos campos das contas, a comparação era verdadeira.

Podemos exibir o conteúdo de ambas:

fmt.Println(contaDaCris)
fmt.Println(contaDaCris2)COPIAR CÓDIGO
No terminal, os valores exibidos serão exatamente os mesmos para ambas, &(Cris 0 0 500}. Entretanto, a comparação é falsa de todo modo. Ele mesmo dá a dica, mostrando o & na frente de cada um dos conjuntos de valores, pois se tratam de endereços diferentes. Vamos pegar apenas o conteúdo desses endereços para comparar.

var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.titular = 500

var contaDaCris2 *ContaCorrente
contaDaCris2 = new(ContaCorrente)
contaDaCris2.titular = "Cris"
contaDaCris2.saldo = 500

fmt.Println(contaDaCris)
fmt.Println(contaDaCris2)

fmt.Println(*contaDaCris == *contaDaCris2)COPIAR CÓDIGO
Agora, mesmo mostrando que os endereços são diferentes, o resultado será que os conteúdos são iguais, com verdadeiro para a comparação no terminal. Poderemos visualizar também o endereço, colocando o & no início de ambas as variáveis na hora de exibi-las.

fmt.Println(&contaDaCris)
fmt.Println(&contaDaCris2)

fmt.Println(*contaDaCris == *contaDaCris2)COPIAR CÓDIGO
Agora será mostrado o local na memória do computador onde a conta estará, um valor complexo, que mistura números e letras, com o final "e020" para contaDaCris, outro com final "e028" para contadaCris2. Sendo "20" e "28" números diferentes, os endereços são diferentes. Se uma casa é a número 20 e a outra número 28, por mais que sejam iguais, tenham o mesmo tamanho e mobília, são casas diferentes.

Só que quando comparamos os conteúdos de ambos os endereços, eles serão iguais. Essas questões servem para aprofundarmos nosso conhecimento quanto aos ponteiros.

Os ponteiros são muito utilizados em linguagens como C e C++ e trabalhar com eles é um grande desafio. O maior benefício de sabermos a diferença entre endereços que alocam informações na memória estará no autodesempenho, ou seja as pessoas que que usam C e C++ conseguem ter um alto desempenho dessa forma, com boa velocidade.

Felizmente no Go vimos que a linguagem tem a capacidade de comparar o conteúdo armazenado em dois endereços independentemente da diferença entre eles.

@@03
Criando o método sacar

Sacar ou depositar um dinheiro em uma conta corrente é um comportamento comum.
Com base no código que desenvolvemos até agora desenvolveremos um comportamento de saque.Removeremos as nossas comparação entre tipos e criaremos uma nova conta corrente, contaDaSilvia.

Ela será do tipo ContaCorrente, terá como titular a string "Silvia" e saldo de 500. Vamos exibir a conta para visualizá-la no nosso terminal.

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia)

}COPIAR CÓDIGO
Executamos e no terminal haverá {Silvia 0 0 500}, ou seja, o nome do titular, dois zeros referentes aos valores de Número da Agência e Conta que não adicionamos, e "500" que é o valor do nosso saldo. Podemos exibir apenas o saldo:

fmt.Println(contaDaSilvia.saldo)

Executaremos mais uma vez e veremos apenas o valor do saldo, "500".

Para sacar um valor, criaremos uma nova variável, valorDoSaque, e usaremos dois pontos igual (:=) para atribuir um valor de "200". Se a conta tem "500" e removeremos "200", ficaremos com "300" de saldo.Então o valor do saldo, contaDaSilvia.saldo será igual a contaDaSilvia.saldo menos valorDoSaque'. Depois, exibiremos mais uma vez o valor que Silvia possui na conta.

valorDoSaque := 200
contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque

fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Quando salvamos, porem, o compilador mostra um alerta de erro, e se passarmos o mouse sobre a linha sublinhada em vermelho entenderemos que o problema é tentamos fazer uma operação entre um valor do tipo float e outro do tipo inteiro. valorDoSaque foi criado como um inteiro, mas o saldo é um float. Nesse caso, bastará colocar um ponto (.) após o valor atribuído ao valorDoSaque e a linguagem entenderá que ele é um float.

valorDoSaque := 200.

Agora, quando executarmos o código, teremos primeiro o "500", valor de saldo que printamos, e depois o cálculo será efetuado e veremos como resultado o valor "300" no terminal. Mas vamos supôr que queremos remover "800" da conta. A cliente não tem esse saldo, então ultrapassaremos o valor.

valorDoSaque := .800

Vamos rodar e o que aparecerá no terminal agora será "-300". Isso não deveria acontecer, pois temos apenas "500" na conta. Para evitar o problema podemos gerar uma verificação para esse saque.

Pensando nisso, criaremos uma função que verificará se o valor que tentaremos sacar será menor do que o valor presente na conta. Vamos criá-la antes da func main(), logo após nossa struct. , chamaremos a função de Sacar() e ela receberá um valor. Nosso valor do saque deverá ser executado nos parênteses da função. Também especificaremos que ele será do tipo float64.

Depois dos parâmetros, retornaremos uma mensagem do tipo string nesse campo, informando se o saque foi realizado com sucesso ou não. Então criaremos uma variável para verificar se é possível sacar com podeSacar :=, pegaremos valorDoSaque e vamos comparar se ele é menor ou igual ao saldo de contaDaSilvia.

Podemos fazer podeSacar := valorDoSaque <= contaDaSilvia.saldo, porém nesse campo precisamos dizer qe quem tentará fazer o saque será a pessoa responsável pela conta. Nesse caso só temos contaDaSilvia, mas se tivermos que adicionar ainda outros clientes, teremos criado uma função que só funciona para contaDaSilvia. Queremos que sempre quando alguém tentar sacar, apontemos para a conta dessa pessoa, de forma semelhante ao que acontece com o this do Java ou o self do Python.

Para referenciarmos esse ponteiro no momento da criação do tipo, podemos colocar entre parênteses logo após func e antes de Sacar() a inscrição (c *ContaCorrente), o que significa que quando a função for chamada, o código apontará para a conta que chama. Nesse caso, quando chamarmos a função, não precisaremos especificar que nos tratamos da conta de um cliente ou outro.

Nesse caso, se a conta que estiver chamando a função tiver saldo, será possível sacar. Criaremos uma condicional if para fazer a verificação. Poderemos sacar se for verdadeiro que valorDoSaque é menor do que saldo. Se podemos sacar colocaremos c.saldo no corpo do if. Podíamos escrever conta.saldo, se escrevêssemos (conta *ContaCorrente) , mas por uma questão da linguagem Go, sempre utilizamos a primeira letra do nome no nosso ponteiro dentro da função.

Porém, escrevemos muito para fazer nossa operação usando contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque. O que queremos fazer é pegar o que há na conta e subtrair pelo valor do saque. Há outra forma de fazer o mesmo, e a utilizaremos na função, com c.saldo -= valorDoSaque.

O valorDoSaque será o valor que passaremos para esse método. Assim, será possível remover o valor para sacar se ele for menor do que o saldo, e retornaremos uma mensagem do tipo string, entre aspas duplas, "Saque realizado com sucesso".

Mas e se isso não for verdade? Então, teremos o else para o caso do valor que tentarmos sacar ser maior do que o saldo. Retornaremos "Saldo insuficiente".

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque <= c.saldo
    if podeSacar {
            c.saldo -= valorDoSaque
            return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}COPIAR CÓDIGO
Criamos função para verificar se o valor do saque é menor do que o do saldo. Se sim, poderemos sacar.

Apagaremos contaDoGustado do fim do código, pois nosso interese nesse momento é visualizar apenas a conta da Silvia. O valor do saldo dela, "500", já estará sendo exibido com um print.

Agora, no lugar de fazer toda a operação em que o saldo é igual ao saldo menos o saque e vamos apenas chamar a função Sacar() .Bastará digitarmos contaDaSilvia e um ponto (.) para que Sacar() seja sugeridp.

Como essa função devolve uma string, podemos colocá-la dentro da string dos parâmetros do print. Quando digitamos contadaSilvia., conseguimos visualizar todos os campos que temos pressionando "Ctrl + espaço" para ter acesso ao atalho. Após contadaSilvia.sacar terá que vir o abrir e fechar de parênteses (), pois queremos executar. Colocando os parênteses, já será recomendado definir o valor do saque. Colocaremos "300".

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia.saldo)

    fmt.Println(contaDaSilvia.Sacar(300))
}COPIAR CÓDIGO
Abriremos o terminal com "Command + J", o limparemos com "Ctrl + L" e executaremos mais uma vez.

Veremos o valor "500" que havia na conta e em seguida a mensagem de que o saque foi realizado com sucesso. Pediremos para que o código nos mostre quanto restou na conta da Silvia depois do saque colocando outro print embaixo do saque.

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(300))
fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Agora veremos que tínhamos "500", realizamos com sucesso o saque (de 300) e será impresso que ficamos com "200". Então nossa função terá funcionado para quando temos saldo. Mas e quando não temos?

Tentaremos sacar "600" para testar.

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(600))
fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Digitaremos "Ctrl + L" e executaremos mais uma vez o programa. Agora veremos a mensagem "Saldo insuficiente" e o valor do saldo permanecerá em "500". Parece que tudo está funcionando, exceto por um detalhe.

Tínhamos "500" de saldo. Ao tentar sacar "600", não foi possível porque "600" é maior do que "500". Só que se o valor do saque for um valor negativo?

Tentaremos sacar "-100".

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(-100))
fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Quanto teremos de saldo quando executarmos a função? "-100" é menor do que "500", condição para que o saque seja realizado.

Quando rodarmos, veremos que a mensagem é a de que o saque de "-100" foi realizado com sucesso, e o valor restante será "600". Ou seja, a operação efetuada foi a do valor do saldo menos menos (- (-)) o valorDoSaque , de forma que o valor ficou positivo e foi atribuído mais "100" ao que havia na conta.

Então, além de verificar se o valor do saque é maior do que o saldo, precisamos simultaneamente verificar se ele é um valor positivo, maior do que 0. Para isso, colocaremos && e assim para que podeSacar, que é um booleano (pode ser verdadeiro ou falso), seja verdadeiro, estas duas condições precisam ser verdadeiras.

podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

Vamos salvar e executar mais uma vez. Veremos que p saldo era "500", tentando sacar "-100" o saldo será insuficiente e serão mantidos os "500" na conta de Silvia. Não foi possível realizar o saque.

Agora se tentarmos sacar outro valor, como "400", o saque será realizado com scesso e ainda haverá "100" na conta.

type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia.saldo)

    fmt.Println(contaDaSilvia.Sacar(400))
    fmt.Println(contaDaSilvia.saldo)

}

@@03
Criando o método sacar

Sacar ou depositar um dinheiro em uma conta corrente é um comportamento comum.
Com base no código que desenvolvemos até agora desenvolveremos um comportamento de saque.Removeremos as nossas comparação entre tipos e criaremos uma nova conta corrente, contaDaSilvia.

Ela será do tipo ContaCorrente, terá como titular a string "Silvia" e saldo de 500. Vamos exibir a conta para visualizá-la no nosso terminal.

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia)

}COPIAR CÓDIGO
Executamos e no terminal haverá {Silvia 0 0 500}, ou seja, o nome do titular, dois zeros referentes aos valores de Número da Agência e Conta que não adicionamos, e "500" que é o valor do nosso saldo. Podemos exibir apenas o saldo:

fmt.Println(contaDaSilvia.saldo)

Executaremos mais uma vez e veremos apenas o valor do saldo, "500".

Para sacar um valor, criaremos uma nova variável, valorDoSaque, e usaremos dois pontos igual (:=) para atribuir um valor de "200". Se a conta tem "500" e removeremos "200", ficaremos com "300" de saldo.Então o valor do saldo, contaDaSilvia.saldo será igual a contaDaSilvia.saldo menos valorDoSaque'. Depois, exibiremos mais uma vez o valor que Silvia possui na conta.

valorDoSaque := 200
contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque

fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Quando salvamos, porem, o compilador mostra um alerta de erro, e se passarmos o mouse sobre a linha sublinhada em vermelho entenderemos que o problema é tentamos fazer uma operação entre um valor do tipo float e outro do tipo inteiro. valorDoSaque foi criado como um inteiro, mas o saldo é um float. Nesse caso, bastará colocar um ponto (.) após o valor atribuído ao valorDoSaque e a linguagem entenderá que ele é um float.

valorDoSaque := 200.

Agora, quando executarmos o código, teremos primeiro o "500", valor de saldo que printamos, e depois o cálculo será efetuado e veremos como resultado o valor "300" no terminal. Mas vamos supôr que queremos remover "800" da conta. A cliente não tem esse saldo, então ultrapassaremos o valor.

valorDoSaque := .800

Vamos rodar e o que aparecerá no terminal agora será "-300". Isso não deveria acontecer, pois temos apenas "500" na conta. Para evitar o problema podemos gerar uma verificação para esse saque.

Pensando nisso, criaremos uma função que verificará se o valor que tentaremos sacar será menor do que o valor presente na conta. Vamos criá-la antes da func main(), logo após nossa struct. , chamaremos a função de Sacar() e ela receberá um valor. Nosso valor do saque deverá ser executado nos parênteses da função. Também especificaremos que ele será do tipo float64.

Depois dos parâmetros, retornaremos uma mensagem do tipo string nesse campo, informando se o saque foi realizado com sucesso ou não. Então criaremos uma variável para verificar se é possível sacar com podeSacar :=, pegaremos valorDoSaque e vamos comparar se ele é menor ou igual ao saldo de contaDaSilvia.

Podemos fazer podeSacar := valorDoSaque <= contaDaSilvia.saldo, porém nesse campo precisamos dizer qe quem tentará fazer o saque será a pessoa responsável pela conta. Nesse caso só temos contaDaSilvia, mas se tivermos que adicionar ainda outros clientes, teremos criado uma função que só funciona para contaDaSilvia. Queremos que sempre quando alguém tentar sacar, apontemos para a conta dessa pessoa, de forma semelhante ao que acontece com o this do Java ou o self do Python.

Para referenciarmos esse ponteiro no momento da criação do tipo, podemos colocar entre parênteses logo após func e antes de Sacar() a inscrição (c *ContaCorrente), o que significa que quando a função for chamada, o código apontará para a conta que chama. Nesse caso, quando chamarmos a função, não precisaremos especificar que nos tratamos da conta de um cliente ou outro.

Nesse caso, se a conta que estiver chamando a função tiver saldo, será possível sacar. Criaremos uma condicional if para fazer a verificação. Poderemos sacar se for verdadeiro que valorDoSaque é menor do que saldo. Se podemos sacar colocaremos c.saldo no corpo do if. Podíamos escrever conta.saldo, se escrevêssemos (conta *ContaCorrente) , mas por uma questão da linguagem Go, sempre utilizamos a primeira letra do nome no nosso ponteiro dentro da função.

Porém, escrevemos muito para fazer nossa operação usando contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque. O que queremos fazer é pegar o que há na conta e subtrair pelo valor do saque. Há outra forma de fazer o mesmo, e a utilizaremos na função, com c.saldo -= valorDoSaque.

O valorDoSaque será o valor que passaremos para esse método. Assim, será possível remover o valor para sacar se ele for menor do que o saldo, e retornaremos uma mensagem do tipo string, entre aspas duplas, "Saque realizado com sucesso".

Mas e se isso não for verdade? Então, teremos o else para o caso do valor que tentarmos sacar ser maior do que o saldo. Retornaremos "Saldo insuficiente".

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque <= c.saldo
    if podeSacar {
            c.saldo -= valorDoSaque
            return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}COPIAR CÓDIGO
Criamos função para verificar se o valor do saque é menor do que o do saldo. Se sim, poderemos sacar.

Apagaremos contaDoGustado do fim do código, pois nosso interese nesse momento é visualizar apenas a conta da Silvia. O valor do saldo dela, "500", já estará sendo exibido com um print.

Agora, no lugar de fazer toda a operação em que o saldo é igual ao saldo menos o saque e vamos apenas chamar a função Sacar() .Bastará digitarmos contaDaSilvia e um ponto (.) para que Sacar() seja sugeridp.

Como essa função devolve uma string, podemos colocá-la dentro da string dos parâmetros do print. Quando digitamos contadaSilvia., conseguimos visualizar todos os campos que temos pressionando "Ctrl + espaço" para ter acesso ao atalho. Após contadaSilvia.sacar terá que vir o abrir e fechar de parênteses (), pois queremos executar. Colocando os parênteses, já será recomendado definir o valor do saque. Colocaremos "300".

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia.saldo)

    fmt.Println(contaDaSilvia.Sacar(300))
}COPIAR CÓDIGO
Abriremos o terminal com "Command + J", o limparemos com "Ctrl + L" e executaremos mais uma vez.

Veremos o valor "500" que havia na conta e em seguida a mensagem de que o saque foi realizado com sucesso. Pediremos para que o código nos mostre quanto restou na conta da Silvia depois do saque colocando outro print embaixo do saque.

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(300))
fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Agora veremos que tínhamos "500", realizamos com sucesso o saque (de 300) e será impresso que ficamos com "200". Então nossa função terá funcionado para quando temos saldo. Mas e quando não temos?

Tentaremos sacar "600" para testar.

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(600))
fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Digitaremos "Ctrl + L" e executaremos mais uma vez o programa. Agora veremos a mensagem "Saldo insuficiente" e o valor do saldo permanecerá em "500". Parece que tudo está funcionando, exceto por um detalhe.

Tínhamos "500" de saldo. Ao tentar sacar "600", não foi possível porque "600" é maior do que "500". Só que se o valor do saque for um valor negativo?

Tentaremos sacar "-100".

fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(-100))
fmt.Println(contaDaSilvia.saldo)COPIAR CÓDIGO
Quanto teremos de saldo quando executarmos a função? "-100" é menor do que "500", condição para que o saque seja realizado.

Quando rodarmos, veremos que a mensagem é a de que o saque de "-100" foi realizado com sucesso, e o valor restante será "600". Ou seja, a operação efetuada foi a do valor do saldo menos menos (- (-)) o valorDoSaque , de forma que o valor ficou positivo e foi atribuído mais "100" ao que havia na conta.

Então, além de verificar se o valor do saque é maior do que o saldo, precisamos simultaneamente verificar se ele é um valor positivo, maior do que 0. Para isso, colocaremos && e assim para que podeSacar, que é um booleano (pode ser verdadeiro ou falso), seja verdadeiro, estas duas condições precisam ser verdadeiras.

podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

Vamos salvar e executar mais uma vez. Veremos que p saldo era "500", tentando sacar "-100" o saldo será insuficiente e serão mantidos os "500" na conta de Silvia. Não foi possível realizar o saque.

Agora se tentarmos sacar outro valor, como "400", o saque será realizado com scesso e ainda haverá "100" na conta.

type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia.saldo)

    fmt.Println(contaDaSilvia.Sacar(400))
    fmt.Println(contaDaSilvia.saldo)

}COPIAR CÓDIGO

@@04
Depositando dez reais

Para criar um método capaz de depositar R$10,00 no saldo de uma conta, foi desenvolvido o seguinte código:
package main

import (
    "fmt"
)

type Conta struct {
    saldo float64
}

func (c *Conta) depositarDezReais() float64 {
    return c.saldo + 10
}

func main() {
    contaTeste := Conta{saldo: 10}

    contaTeste.depositarDezReais()
    contaTeste.depositarDezReais()

    fmt.Println(contaTeste)
}COPIAR CÓDIGO
Com base no trecho de código acima, marque as alternativas verdadeiras.

Ao executar o código, uma mensagem de erro informará que o saldo está incorreto.
 
O código acima possui um erro de lógica e não de sintaxe. Sendo assim, nenhuma mensagem de erro será exibida ao executar.
Alternativa correta
Ao executar o código acima, teremos o saldo de 20 como resultado.
 
Ao executar esse código, não teremos como resultado do saldo o valor de 20.
Alternativa correta
Ao executar o código acima, teremos o saldo de 30 como resultado.
 
Ao executar esse código, não teremos como resultado do saldo o valor de 30.
Alternativa correta
Ao executar o código acima, teremos o saldo de 10 como resultado.
 
Certo! Como não atribuímos o código no saldo da conta( c.saldo = c.saldo + 10), nosso saldo continuará sendo 10. Neste link, você pode executar o código acima código.

https://play.golang.org/p/aG6sxPAuk2t

@@05
Para saber mais

Função com quantidade de parâmetros indeterminados
Conforme estudado nesta aula, uma função em Go pode ter um, muitos ou nenhum parâmetro:

package main

import (
    "fmt"
)

func SemParametro() string {
    return "Exemplo de função sem parâmetro!"
}

func UmParametro(texto string) string {
    return texto
}

func DoisParametros(texto string, numero int) (string, int) {
    return texto, numero
}

func main() {
    fmt.Println(SemParametro())
    fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
    fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))
}
COPIAR CÓDIGO
E nossa saída seria:



Neste link, você pode executar o código acima.

Porém, é possível que uma função receba um número indeterminado de parâmetros. Funções deste tipo são conhecidas em Go como variadic functions, ou função variádica.

Criando um função variádica
Para criar uma variadic function, devemos preceder o tipo do argumento com reticências, conforme o exemplo abaixo:

package main

import (
    "fmt"
)

func Somando(numeros ...int) int {
    resultadoDaSoma := 0
    for _, numero := range numeros {
        resultadoDaSoma += numero
    }
    return resultadoDaSoma
}

func main() {
    fmt.Println(Somando(1))
    fmt.Println(Somando(1,1))
    fmt.Println(Somando(1,1,1))
    fmt.Println(Somando(1,1,2,4))
}COPIAR CÓDIGO
Neste link, você pode executar o código acima.

Note o uso das reticências na declaração do parâmetro número: numeros ...int. Portanto, podemos criar uma função sem parâmetro, com um, dois, três, ou uma quantidade indeterminada de parâmetros com Go.

https://play.golang.org/p/Kq5ZXTXLVrW

https://en.wikipedia.org/wiki/Variadic_function

https://play.golang.org/p/s-kHUpafjOP

@@06
Faça como eu fiz na aula

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito, excelente. Se ainda não, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código aqui escrito.
Se por acaso você já domina essa parte, em cada capítulo, você poderá baixar o projeto feito até aquele ponto.

O gabarito deste exercício é o passo a passo demonstrado no vídeo. Tenha certeza de que tudo está certo antes de continuar. Ficou com dúvida? Podemos te ajudar pelo nosso fórum.

@@07
O que aprendemos?

Nessa aula:
Criamos um nova conta corrente utilizando a palavra new;
Em seguida, comparamos os tipos criados comparando suas referências e entendemos o que são ponteiros na prática;
Para finalizar, desenvolvemos o método sacar que verifica se o valor do saque é maior do que zero e se a conta possui saldo.
Projeto desenvolvido nesta aula
Neste link, você fará o download do projeto feito até esta aula.

Caso queira visualizar o código desenvolvido até aqui, clique neste link.

Na próxima aula
Vamos criar um método que possibilita o depósito, outro método para realizar a transferência de dinheiro entre contas e dividir nosso código em pacotes, trabalhando com visibilidade dos campos!

https://github.com/alura-cursos/go_oo/archive/aula2.zip

https://github.com/alura-cursos/go_oo/tree/aula2

