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