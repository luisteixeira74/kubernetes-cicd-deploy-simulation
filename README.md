# kubernetes-cicd-deploy-simulation

Este projeto simula um pipeline de CI/CD com Docker e Kubernetes, usando uma API simples em Go como exemplo de aplicação.

## 🚀 Etapa 1: Build e Teste Local com Docker

Antes de aplicar ao Kubernetes, valide a aplicação localmente:

```bash
# Build da imagem Docker local
docker build -t go-k8s-cicd-example .

# Execute o container localmente
docker run -p 8080:8080 go-k8s-cicd-example
```

Acesse em: [http://localhost:8080](http://localhost:8080)  
Você verá a mensagem: `Hello, Kubernetes! 👋`

---

## 🧪 Etapa 2: Criar o Cluster Kubernetes com Kind

Caso ainda não tenha um cluster local:

```bash
kind create cluster --name k8s-simulation
```

Isso criará um cluster Kubernetes local chamado `k8s-simulation`.

---

## 📦 Etapa 3: Enviar a Imagem Docker para o Cluster Kind

Por padrão, o Kind **não tem acesso direto** às suas imagens Docker locais.

Para que o cluster consiga usar a imagem `go-k8s-cicd-example`, execute:

```bash
kind load docker-image go-k8s-cicd-example:latest --name k8s-simulation
```

---

## ☸️ Etapa 4: Deploy no Kubernetes

Com o cluster criado e a imagem carregada, aplique os manifests:

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

Verifique os pods:

```bash
kubectl get pods
```

O pod deve estar com STATUS `Running`.

---

## 🧹 Dica de Debug (se necessário)

Se o pod estiver com `ErrImagePull`, isso indica que o Kind não conseguiu acessar sua imagem local.

Solução:

```bash
kind load docker-image go-k8s-cicd-example:latest --name k8s-simulation
kubectl delete pod -l app=go-k8s-app  # ou aguarde a recriação automática
```

---

## ✅ Acesso Externo (porta NodePort)

Se você configurou o `type: NodePort` no `service.yaml`, acesse a aplicação via:

```bash
curl http://localhost:30000
```

Você deve ver: `Hello, Kubernetes! 👋` 

---

## 📁 Estrutura

```
.
├── main.go
├── go.mod
├── Dockerfile
├── k8s/
│   ├── deployment.yaml
│   └── service.yaml
└── README.md
```

---

## 🛠️ Requisitos

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Kind](https://kind.sigs.k8s.io/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)