# Implementação do Algoritmo k-Nearest Neighbors (k-NN) para Predição de Churn

## Descrição

Este projeto tem como objetivo implementar manualmente o algoritmo de classificação **k-Nearest Neighbors (k-NN)** utilizando a linguagem Go, sem o uso de bibliotecas de Machine Learning.

A implementação contempla todas as etapas fundamentais do processo de classificação:

* Leitura de dados a partir de arquivos CSV;
* Pré-processamento dos dados;
* Codificação de atributos categóricos utilizando One-Hot Encoding;
* Normalização Min-Max dos atributos contínuos;
* Divisão do dataset em conjuntos de treino e teste;
* Implementação manual do algoritmo k-NN;
* Avaliação do modelo utilizando acurácia;
* Comparação dos resultados para diferentes valores de k.

O dataset utilizado é composto por informações de clientes e possui como objetivo prever a ocorrência de churn.

---

# Objetivos do Trabalho

Dado um conjunto de clientes contendo atributos demográficos e comportamentais, o sistema deverá classificar se um cliente possui churn ou não utilizando o algoritmo k-NN.

Classe alvo:

| Valor | Significado       |
| ----- | ----------------- |
| 0     | Cliente ativo     |
| 1     | Cliente com churn |

---

# Requisitos do Trabalho

## Entrada dos Dados

* [x] Leitura de arquivo CSV
* [x] Conversão das linhas para estruturas Go
* [x] Identificação da coluna de classe (churned)

---

## Pré-processamento

### Identificação dos atributos

* [ ] Identificação dos atributos contínuos
* [ ] Identificação dos atributos categóricos

### Codificação dos atributos categóricos

* [ ] Implementação manual de One-Hot Encoding
* [ ] Conversão de atributos categóricos para representação numérica

### Normalização

* [ ] Implementação manual da normalização Min-Max
* [ ] Aplicação da normalização apenas sobre atributos contínuos

---

## Divisão do Dataset

* [ ] Embaralhamento dos dados
* [ ] Utilização de semente fixa para reprodução dos resultados
* [ ] Divisão 75% treino e 25% teste

---

## Implementação do k-NN

* [ ] Implementação da distância euclidiana
* [ ] Cálculo da distância entre instâncias de teste e treino
* [ ] Seleção dos k vizinhos mais próximos
* [ ] Classificação por votação majoritária

---

## Avaliação

* [ ] Execução para k = 1
* [ ] Execução para k = 3
* [ ] Execução para k = 5
* [ ] Execução para k = 7
* [ ] Execução para k = 9
* [ ] Cálculo da acurácia
* [ ] Comparação dos resultados

---

# Progresso da Implementação

## Estruturas de Domínio

* [x] Estrutura Customer
* [ ] Estrutura Sample
* [ ] Estrutura Neighbor

## Leitura dos Dados

* [x] CSV Reader
* [x] Conversão para Customer

## Pré-processamento

* [ ] Descoberta automática de categorias
* [ ] OneHotEncoder
* [ ] MinMaxScaler
* [ ] Dataset Builder

## Treino e Teste

* [ ] TrainTestSplit

## Algoritmo

* [ ] Distância Euclidiana
* [ ] Classificador KNN
* [ ] Votação Majoritária

## Avaliação

* [ ] Accuracy

## Relatório

* [ ] Tabela de resultados
* [ ] Análise dos valores de k

---

# Classificação dos Atributos

## Atributos Contínuos

| Atributo                 | Tipo     |
| ------------------------ | -------- |
| age                      | Contínuo |
| total_orders             | Contínuo |
| total_spend_usd          | Contínuo |
| avg_order_value_usd      | Contínuo |
| days_since_last_purchase | Contínuo |
| reviews_given            | Contínuo |
| avg_review_score         | Contínuo |
| returns_made             | Contínuo |
| wishlist_items           | Contínuo |

Esses atributos serão normalizados utilizando Min-Max.

---

## Atributos Categóricos

| Atributo                 | Tipo       |
| ------------------------ | ---------- |
| country                  | Categórico |
| gender                   | Categórico |
| membership_tier          | Categórico |
| preferred_category       | Categórico |
| preferred_device         | Categórico |
| preferred_payment_method | Categórico |
| acquisition_channel      | Categórico |
| newsletter_subscribed    | Categórico |

Esses atributos serão codificados utilizando One-Hot Encoding.

---

## Classe Alvo

| Atributo | Tipo   |
| -------- | ------ |
| churned  | Classe |

Valores:

| Valor | Significado      |
| ----- | ---------------- |
| 0     | Não possui churn |
| 1     | Possui churn     |

---

# Fluxo da Aplicação

```text
CSV
 │
 ▼
Leitura dos Dados
 │
 ▼
Customer
 │
 ▼
One-Hot Encoding
 │
 ▼
Normalização Min-Max
 │
 ▼
Sample Dataset
 │
 ▼
Train/Test Split
 │
 ├── 75% Treino
 │
 └── 25% Teste
 │
 ▼
k-NN
 │
 ▼
Predição
 │
 ▼
Acurácia
 │
 ▼
Relatório
```

---

# Valores de K Avaliados

```text
k = 1
k = 3
k = 5
k = 7
k = 9
```