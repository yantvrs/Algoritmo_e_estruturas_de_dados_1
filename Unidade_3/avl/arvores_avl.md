## Árvores balanceadas 
- A eficiência da busca em uma árvore binária depende do seu balanceamento

"Problema"
- Algoritmos de inserção e remoção não garantem que a árvore gerada a cada passo seja balanceada 
- Sequência de inserções em ordem de "escada"

- Custo da inserção, busca e remoção em uma árvore binária 
    - balanceada : O(log(n))
    - Não balanceada : O(n)
"n" corrensponde ao número de nós na árvore

- É uma árvore binária onde as alturas das sub-árvores "esquerda" e "direita" de cada nó diferem de no máximo uma unidade

- Essa diferença é chamada de "fator de balanceamento" do nó, que é calculada através da formula abaixo:
    bf = hr - hl

# Árvores balanceadas (AVL)

    - É um tipo de árvore binária balanceada 
    - Criada por "A"delson-"V"elskii e "L"andis em 1962

    - Permite rebalanceamento local
        - apenas a parte afetada pela "inserção" ou "remoção" é rebalanceada
        - Uso de rotações "simples" ou "duplas" na etapa de rebalanceamento

    - A árvore avl busca mater-se como uma árvore binária quase completa
    - Custo de qualquer algoritmo é máximo O(log(n))

    - "Fator de balanceamento" ou "bf" 
    - Diferença nas alturas das sub-árvores "esquerda" e "direita" 
      - Se uma das sub-árvores não existir, sua altura será -1 


# Tipos de rotação 

"Rotação"
- Operação básica para balanceamento da AVL 
- Dois tipos de rotações "simples" e "duplas"

"Simples"
- O nó desbalanceado e seu filho estão no mesmo sentido da inclinação 

"Duplas"
- O nó está desbalanceado e seu filho estão inclinado no sentido inverso ao pai
- Equivale a duas rotações "simples"

- Existem 2 rotações "simples" e 2 "duplas"
  - rotação a "direita e a "esquerda"

  




  Como balancear uma árvore? Algoritmo AVL

  . Inserção na subárvore direita do filho à direita
    > Solução: Rotação simples à esquerda
  . Inserção na subárvore esquerda do filho à esquerda
    > Solução: Rotação simples à direita
  . Inserção na subárvore esquerda do filho à direita
    > Solução: Rotação dupla direita-esquerda
  . Inserção na subárvore direita do filho à esquerda
    > Solução: Rotação dupla esquerda-direita
  
