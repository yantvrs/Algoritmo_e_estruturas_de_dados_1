# Resumo: Heap

Uma heap é uma estrutura de dados em forma de árvore binária, onde cada nó possui um valor maior ou igual aos seus nós filhos (na heap máxima) ou menor ou igual aos seus nós filhos (na heap mínima). É uma implementação eficiente para filas de prioridade.

A principal propriedade de uma heap é a "propriedade de heap", que garante que o valor de cada nó esteja de acordo com sua prioridade em relação aos nós filhos. Na heap máxima, o nó pai possui um valor maior ou igual aos seus nós filhos. Na heap mínima, o nó pai possui um valor menor ou igual aos seus nós filhos.

A estrutura da heap é mantida através das operações de "inserção" e "remoção". Ao inserir um novo elemento, ele é colocado na posição correta de acordo com sua prioridade. A remoção do elemento de maior prioridade é feita ao trocá-lo com o último elemento da heap e, em seguida, ajustando a posição do novo elemento raiz para preservar a propriedade de heap.

As heaps são frequentemente usadas para implementar filas de prioridade, onde o elemento com a maior (ou menor) prioridade pode ser acessado rapidamente. Essa estrutura de dados é eficiente para operações de inserção e remoção, com complexidade O(log n), onde n é o número de elementos na heap.
