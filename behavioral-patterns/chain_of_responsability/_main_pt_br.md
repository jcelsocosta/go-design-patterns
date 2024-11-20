# Explicação sobre o padrão chain of responsability

## Chain of Responsability
O Chain of Responsibility é um padrão de projeto
comportamental que permite que você passe pedidos por uma
corrente de handlers. Ao receber um pedido, cada handler
decide se processa o pedido ou o passa adiante para o próximo
handler na corrente.[1]

## Quando usar
* Quando você quer evitar acoplamento direto entre o remetente de uma solicitação e os manipuladores.
* Quando você precisa aplicar várias etapas de processamento que podem ser dinâmicas ou configuráveis.
* Quando os objtos que compoẽm a cadeia podem ser modificados em tempo de execução
* Quando você quer evitar que um objeto tenha dependências diretas de todos os manipuladores possíveis.

## Referências bibliográficas
[1] - Megulho nos padrões de projetos, Alexander Shvets  
[2] - Go Design Patterns, Mario Castro Contreras  
