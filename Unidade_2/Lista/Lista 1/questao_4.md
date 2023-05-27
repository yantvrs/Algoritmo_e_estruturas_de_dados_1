# Questão 4) A seguir são apresentados 5 fatos sobre algoritmos de ordenação. Planeje, execute experimentos, e apresente resultados que evidenciem cada afirmação.

A. Para vetores de tamanho pequeno, a performance da maioria dos algoritmos de ordenação não vai influenciar, independente da disposição dos elementos.

    a. Sugestão: um algoritmo O(n²), um algoritmo O(nlogn), e como exceção um algoritmo O(k+n)

Resposta: 
Para verificar a afirmação A, que diz que para vetores de tamanho pequeno a performance da maioria dos algoritmos de ordenação não vai influenciar, independente da disposição dos elementos, podemos realizar experimentos comparando três algoritmos de ordenação: um algoritmo O(n²), um algoritmo O(nlogn) e um algoritmo O(k+n), onde k é um valor constante.

B.Vetor de tamanho grande, a performance do algoritmo influencia de forma significativa. Além disso, dependendo da disposição (e valores) dos elementos no vetor, podemos experimentar performances bem diferentes (melhor e pior caso).
    b.Sugestão: um algoritmo O(n²), um algoritmo O(nlogn), um algoritmo O(k+n)

Resposta: 

Sim, a afirmação é verdadeira. Para vetores de tamanho grande, a performance do algoritmo de ordenação pode influenciar significativamente o tempo de execução. Além disso, dependendo da disposição e dos valores dos elementos no vetor, podemos experimentar diferentes cenários de desempenho, como o melhor caso (quando o vetor já está ordenado ou quase ordenado) e o pior caso (quando o vetor está inversamente ordenado ou possui elementos repetidos).

Os experimentos realizados com os algoritmos Bubble Sort, Merge Sort e Counting Sort nos três cenários (vetor ordenado, vetor inversamente ordenado e vetor aleatório) devem ter demonstrado essa diferença de desempenho entre os casos.

No caso do Bubble Sort, por exemplo, ele possui uma complexidade de tempo O(n²), o que significa que seu tempo de execução aumenta quadraticamente com o tamanho do vetor. Portanto, para vetores de tamanho grande, o Bubble Sort terá um desempenho significativamente pior em comparação com algoritmos de complexidade O(nlogn) como o Merge Sort.

O Counting Sort, por sua vez, possui uma complexidade de tempo O(k+n), onde k é o valor máximo no vetor. Esse algoritmo é mais eficiente em casos específicos, como quando o intervalo de valores no vetor é pequeno. Portanto, dependendo dos valores presentes no vetor, o Counting Sort pode apresentar um desempenho melhor em relação aos outros algoritmos.

Em suma, a escolha do algoritmo de ordenação correto para um vetor de tamanho grande e com diferentes disposições e valores dos elementos é fundamental para obter um desempenho satisfatório.

C. MergeSort tem sempre um desempenho muito bom, independente da disposição dos elementos no vetor.

Resposta: 
A afirmação de que o MergeSort tem sempre um desempenho muito bom, independente da disposição dos elementos no vetor, não é totalmente precisa. Embora o MergeSort seja conhecido por sua eficiência em uma ampla gama de cenários, seu desempenho pode variar dependendo da disposição dos elementos no vetor.

O MergeSort possui uma complexidade de tempo O(nlogn), o que o torna um algoritmo de ordenação eficiente na maioria dos casos. Ele divide o vetor em partes menores, realiza a ordenação dessas partes e, em seguida, combina as partes ordenadas para obter o vetor finalmente ordenado.

No entanto, embora o MergeSort tenha uma complexidade de tempo ideal, é importante observar que ele requer espaço adicional para armazenar as partes divididas do vetor durante o processo de mesclagem. Isso pode se tornar uma desvantagem em casos em que o vetor é extremamente grande e o espaço disponível na memória é limitado.

Além disso, embora o MergeSort possua uma eficiência teórica consistente, sua implementação prática pode ser influenciada pelo uso de memória cache e outros fatores de hardware. Em vetores pequenos ou com um tamanho específico que corresponde ao tamanho da memória cache, o MergeSort pode apresentar um desempenho relativamente pior em comparação com algoritmos de ordenação mais simples, como o Insertion Sort ou o Bubble Sort.

Portanto, embora o MergeSort seja geralmente considerado um algoritmo eficiente e com bom desempenho, é necessário considerar as limitações práticas, como o uso de memória e fatores específicos de hardware, bem como a disposição dos elementos no vetor, para avaliar seu desempenho em diferentes cenários.

D. O pior caso do Quicksort é com o vetor ordenado de forma crescente/decrescente. O Quicksort com randomização de pivô resolve esse mau desempenho.

A afirmação de que o pior caso do Quicksort ocorre quando o vetor está ordenado de forma crescente ou decrescente é verdadeira. No pior caso, o Quicksort tem uma complexidade de tempo de O(n^2), onde n é o tamanho do vetor. Isso ocorre quando o pivô escolhido sempre resulta em partições desequilibradas, ou seja, uma das partições tem tamanho próximo a n-1 e a outra tem tamanho 0.

No entanto, a afirmação de que o Quicksort com randomização de pivô resolve esse mau desempenho também é verdadeira. A randomização do pivô é uma técnica que envolve escolher aleatoriamente um elemento do vetor como pivô durante a execução do algoritmo. Essa abordagem reduz significativamente a probabilidade de ocorrer o pior caso mencionado acima.

Ao escolher aleatoriamente o pivô, as chances de obter partições equilibradas aumentam consideravelmente. Isso ocorre porque o pivô é selecionado aleatoriamente de forma independente da ordem dos elementos no vetor. Portanto, mesmo que o vetor esteja ordenado de forma crescente ou decrescente, a randomização do pivô tende a criar partições mais equilibradas, melhorando o desempenho do Quicksort.

Com a randomização do pivô, o Quicksort tem uma complexidade média de tempo de O(nlogn), tornando-o um dos algoritmos de ordenação mais eficientes. No entanto, é importante ressaltar que, embora a randomização do pivô minimize a probabilidade do pior caso, ainda existe uma pequena chance de ocorrer um desempenho abaixo do ideal em cenários específicos.

Em resumo, a randomização do pivô no Quicksort é uma técnica eficaz para evitar o pior caso em vetores ordenados de forma crescente ou decrescente. Isso resulta em um desempenho médio muito melhor, tornando o Quicksort uma opção poderosa para ordenação eficiente de vetores em geral.

E. Explique quando o CountingSort tem bom desempenho e quando tem mau desempenho mostrando os resultados através dos experimentos.

Resposta: 
A afirmação é verdadeira. O desempenho do Counting Sort pode variar dependendo das características do vetor de entrada. Vamos analisar quando o Counting Sort tem bom desempenho e quando tem mau desempenho.

O Counting Sort é um algoritmo de ordenação eficiente quando o intervalo de valores no vetor de entrada é relativamente pequeno em comparação ao tamanho do vetor. Isso ocorre porque o Counting Sort utiliza uma contagem de ocorrências dos valores presentes no vetor para realizar a ordenação.

Quando o vetor de entrada possui um intervalo de valores pequeno, o Counting Sort pode executar em tempo linear O(k+n), onde k é o valor máximo no vetor e n é o tamanho do vetor. Nesse caso, o Counting Sort tem um desempenho excelente, já que sua complexidade de tempo é linear, independentemente da distribuição dos valores no vetor.

Por outro lado, quando o intervalo de valores no vetor de entrada é muito grande em relação ao tamanho do vetor, o Counting Sort se torna ineficiente. Isso ocorre porque o Counting Sort requer uma contagem para cada valor único no vetor, o que pode exigir uma quantidade significativa de espaço de memória.

Além disso, quando o vetor de entrada contém muitas repetições de valores, o Counting Sort também pode ter um desempenho inferior. Isso ocorre porque a contagem dessas repetições pode levar a um aumento no tempo e espaço necessário para executar o algoritmo.

Para demonstrar essas situações de bom e mau desempenho do Counting Sort, você pode executar experimentos com diferentes vetores de entrada. Considere vetores com diferentes intervalos de valores e repetições. A partir desses experimentos, você poderá observar que o Counting Sort tem um desempenho muito bom quando o intervalo de valores é pequeno e mau desempenho quando o intervalo de valores é grande ou há muitas repetições.

Lembre-se de medir o tempo de execução do Counting Sort em cada cenário e comparar os resultados para avaliar o desempenho.