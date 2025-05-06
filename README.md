# go-k8s-cicd-example

Este projeto demonstra um pipeline CI/CD simples usando GitHub Actions, Docker Hub e um cluster Kubernetes local (via kind ou minikube). A aplicação é um serviço HTTP simples escrito em Go, com deploy automatizado da imagem no Docker Hub.

## 📦 Requisitos

- **Go (Golang)** — aplicação principal
- **Docker** — criação de imagem
- **Docker Hub** — armazenamento da imagem
- **GitHub Actions** — pipeline CI/CD automatizado
- **Kubernetes (Kind)** — cluster local para deploy
- **kubectl** — para gerenciar o cluster local

## 🚀 Setup local

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/go-k8s-cicd-example.git
cd go-k8s-cicd-example
```

2. Construa a imagem localmente (opcional):

```bash
docker build -t lfmacedo/go-k8s-cicd-example:latest .
```

3. Suba o cluster com Minikube (exemplo):

```bash
minikube start
kubectl apply -f k8s/deployment.yaml
```

4. Teste o serviço:

```bash
curl http://localhost:8080/hello
```

## ☁️ GitHub Actions (CI/CD)

O workflow está em `.github/workflows/docker-deploy.yml`.

### O que faz:

- Build da imagem Docker.
- Push da imagem para Docker Hub.
- ⚠️ **Não realiza o deploy no cluster automaticamente**. Isso ocorre porque o GitHub Actions não tem acesso ao seu cluster local.

### Como tratar isso:

- O deploy no cluster deve ser feito manualmente com:

```bash
kubectl set image deployment/go-k8s-app go-k8s-container=lfmacedo/go-k8s-cicd-example:latest
kubectl rollout status deployment/go-k8s-app
```

## 🔐 Secrets necessários no GitHub

Crie estes secrets no GitHub (Settings > Secrets and variables > Actions):

| Nome                | Descrição                          |
|---------------------|--------------------------------------|
| `DOCKER_USERNAME`   | Seu nome de usuário Docker Hub       |
| `DOCKER_PASSWORD`   | Senha da sua conta Docker Hub        |

## 📁 Estrutura do projeto

```
.
├── Dockerfile
├── main.go
├── k8s
│ ├── deployment.yaml # Deployment + Service com imagem final
│ └── service-nodeport.yaml # Exposição do serviço via NodePort (padrão opcional)
└── .github
└── workflows
└── docker-deploy.yml # Pipeline de CI/CD (sem deploy automático)
```

## ✅ Testando a imagem publicada

```bash
docker pull lfmacedo/go-k8s-cicd-example:latest
docker run -p 8080:8080 lfmacedo/go-k8s-cicd-example:latest
curl http://localhost:8080/hello
```

---

✅ Deploy Manual com kubectl
Suba o cluster local (caso use Kind):

kind create cluster --name go-cicd
Aplique os manifests do Kubernetes:

kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service-nodeport.yaml

Verifique os pods:
kubectl get pods

Verifique o serviço:
kubectl get svc

Acesse a aplicação localmente (porta 8080):

curl http://localhost:8080/hello

Deve retornar:
Hello, World!