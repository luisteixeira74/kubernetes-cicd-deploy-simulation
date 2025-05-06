# go-k8s-cicd-example

![CI/CD](https://github.com/luisteixeira74/kubernetes-cicd-deploy-simulation/actions/workflows/docker-deploy.yml/badge.svg)

Este projeto demonstra um pipeline CI/CD simples usando **GitHub Actions**, **Docker Hub** e um cluster **Kubernetes local** (via Kind ou Minikube).  
A aplicação é um serviço HTTP simples escrito em Go, com deploy automatizado da imagem no Docker Hub.

---

## 📦 Requisitos

- **Go (Golang)** — aplicação principal
- **Docker** — criação da imagem
- **Docker Hub** — armazenamento da imagem
- **GitHub Actions** — pipeline CI/CD automatizado
- **Kubernetes (Kind ou Minikube)** — cluster local para deploy
- **kubectl** — para gerenciar o cluster

---

## 🚀 Setup local

1. Clone o repositório:

```bash
git clone https://github.com/luisteixeira74/kubernetes-cicd-deploy-simulation.git
cd kubernetes-cicd-deploy-simulation
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

---

## ☁️ GitHub Actions (CI/CD)

O workflow está em:  
`.github/workflows/docker-deploy.yml`

### O que ele faz:

- ✅ Faz o build da imagem Docker
- ✅ Faz o push da imagem para o Docker Hub
- ⚠️ **Não realiza o deploy no cluster automaticamente**  
  (isso ocorre porque o GitHub Actions não tem acesso ao seu cluster local)

### Como realizar o deploy:

Manual, com:

```bash
kubectl set image deployment/go-k8s-app go-k8s-container=lfmacedo/go-k8s-cicd-example:latest
kubectl rollout status deployment/go-k8s-app
```

---

## 🔐 Secrets necessários no GitHub

Crie os seguintes *secrets* em:

`Settings > Secrets and variables > Actions`

| Nome                | Descrição                          |
|---------------------|-----------------------------------|
| `DOCKER_USERNAME`   | Seu nome de usuário no Docker Hub |
| `DOCKER_PASSWORD`   | Senha da sua conta Docker Hub     |

---

## 📁 Estrutura do projeto

```
.
├── Dockerfile
├── main.go
├── app/
│   ├── handler.go            # Lógica da API
│   └── handler_test.go       # Teste da API (GET /hello)
├── k8s/
│   ├── deployment.yaml       # Deployment + Service
│   └── service-nodeport.yaml # Exposição via NodePort
├── .github/
│   └── workflows/
│       └── docker-deploy.yml # CI/CD pipeline

```

---

## ✅ Testando a imagem publicada

```bash
docker pull lfmacedo/go-k8s-cicd-example:latest
docker run -p 8080:8080 lfmacedo/go-k8s-cicd-example:latest
curl http://localhost:8080/hello
```

---

## ☸️ Deploy Manual com kubectl (Kind)

1. Suba o cluster local com Kind:

```bash
kind create cluster --name go-cicd
```

2. Aplique os manifests do Kubernetes:

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service-nodeport.yaml
```

3. Verifique os pods:

```bash
kubectl get pods
```

4. Verifique o serviço:

```bash
kubectl get svc
```

5. Acesse a aplicação localmente (porta 8080):

```bash
curl http://localhost:8080/hello
```

**Retorno esperado:**
```
Hello, World!
```