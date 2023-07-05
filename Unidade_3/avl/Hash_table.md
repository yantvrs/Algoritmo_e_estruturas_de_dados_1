# Tabelas hash

"Busca" 
    - Muitos métodos de busca funcinam seguindo o mesmo pricípio: procurar a informação desejada com base na comparação de suas chaves, isto é com base em algum valor que a compõe.

"Problema"
    - Algoritmos eficientes exigem que os elementos sejam armazenados de forma ordenada.
    - Custo da ordenação no melhor caso: "O(nlog(n))"
    - Custo de busca: "O(log(n))"

"Busca ideal"
    - Acesso direto ao elemento procurado, sem nenhuma etapa de comparação de chaves: custo "O(1)"

"Arrays"
    - Estruturas que utilizam "índices" para armazenar informções
    - Permite acessar uma determinada posição do array com custo "O(1)"

"Problema"
    - "Arrays" não possuem nenhum mecanismo que permita calcular a posição onde uma informação está armazenada, de modo que a operação de busca em um array não é "O(1)"

"SOlução"
    - Usar uma "tabela hash"

"Tabela Hash"
    - Também conhecidas como tabelas de "indexação" ou de "espalhamento"
    - É uma generalização da ideia de "array"
    - Utiliza uma função para espalhar os elementos que queremos armazenar na tabela
    - O espalhamento faz com que os elementos fiquem dispersos de forma não ordenada dentro do "array" que define a tabela

"Função de hashing"
    - Função que espalha os elementos na tabela

" Por que espalhar os elementos melhora a busca?
    - Uma tabela hash permite a associação de "valores" a "chaves"
    - "chave": parte da informação que compõe o elemento a ser inserido ou buscado na tabela 
    - "valor" : é a posição (índice) onde o elemento se encontra no "array" que define a tabela

    - Assim, a partir de uma "chave" podemos acessar de forma rápida uma determinada "posição" do "array"
    - Na média, essa operação tem custo "O(1)"

"Vantagens"
    - alta eficiência na operação de busca
    - tempo de busca é praticamente independente do número de "chaves" armazenadas na tabela 
    - implementação simples

"Desvantagens"
    - alto custo para recuperar os elementos da tabela ordenados pela "chave". Nesse caso, é preciso ordenar a tabela 
    - pior caso é "O(n)", sendo "n" o tamanho da tabela: alto número de "colisões"

"Aplicações"
    - busca de elementos em base de dados
    - criptografia: MDS e família SHA(Secure Hash Algorithm)
    - implementação da tabela de símbolos dos compiladores
    - armazenamento de senhas com segurança: a senha não é armazenada no servidor, mas sim o resultado da função hash
    - verificação de integridade de dados e autenticação de mensagens

" Tabela Hash"
    - Sua implementação utiliza uma estrutura similar a da "Lista Sequencial Estática"
    - Utiliza uma "array para armazenar os elementos

"Desvantagem"
    - Necessitar que se defina o tamanho do array previamente.
    - Isso limita o número de elementos que podemos armazenar.

