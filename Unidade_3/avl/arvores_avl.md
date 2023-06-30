# Árvores balanceadas (AVL)

- É uma árvore binária onde as alturas das sub-árvores "esquerda" e "direita" de cada nó diferem de no máximo uma unidade

- Essa diferença é chamada de "fator de balanceamento" do nó, que é calculada através da formula abaixo:
    bf = hr - hl

## Árvores balanceadas 
- A eficiência da busca em uma árvore binária depende do seu balanceamento

"Problema"
- Algoritmos de inserção e remoção não garantem que a árvore gerada a cada passo seja balanceada 
- Sequência de inserções em ordem de "escada"

- Custo da inserção, busca e remoção em uma árvore binária 
    - balanceada : O(log(n))
    - Não balanceada : O(n)
"n" corrensponde ao número de nós na árvore

 
  Como balancear uma árvore? Algoritmo AVL

  . Inserção na subárvore direita do filho à direita
    > Solução: Rotação simples à esquerda
  . Inserção na subárvore esquerda do filho à esquerda
    > Solução: Rotação simples à direita
  . Inserção na subárvore esquerda do filho à direita
    > Solução: Rotação dupla direita-esquerda
  . Inserção na subárvore direita do filho à esquerda
    > Solução: Rotação dupla esquerda-direita
  
