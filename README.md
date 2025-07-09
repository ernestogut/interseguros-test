# Interseguros Test

Este proyecto contiene dos aplicaciones principales: una API en Go (Fiber) y una API en Node.js (Express), junto con infraestructura como código usando Terraform para AWS.

## Estructura del Proyecto

- `express-app/`: API en Node.js (Express)
- `fiber-app/`: API en Go (Fiber)
- `infra/`: Archivos Terraform para infraestructura en AWS
- `docker-compose.yaml`: Orquestación de contenedores para desarrollo local
- `scripts/`: Scripts útiles para build y push de imágenes

---

## Levantar el Proyecto en Local con Docker

### Requisitos

- Docker y Docker Compose instalados

### Pasos

1. Clona el repositorio:
   ```bash
   git clone <url-del-repo>
   cd interseguros-test
   ```
2. Levanta los servicios:
   ```bash
   docker-compose up --build
   ```
3. Accede a las aplicaciones:
   - Fiber (Go): [http://localhost:8080](http://localhost:8080)
   - Express (Node.js): [http://localhost:3000](http://localhost:3000)

### Pruebas Locales

- **Express:**
  ```bash
  docker-compose exec express-app npm test
  ```
- **Fiber:**
  Ejecuta pruebas unitarias dentro del contenedor:
  ```bash
  docker-compose exec fiber-app go test ./...
  ```

---

## Despliegue en la Nube (AWS) con Terraform

### Requisitos

- AWS CLI configurado
- Terraform instalado

### Pasos

1. Entra a la carpeta de infraestructura:
   ```bash
   cd infra
   ```
2. Inicializa Terraform:
   ```bash
   terraform init
   ```
3. Revisa el plan de despliegue:
   ```bash
   terraform plan
   ```
4. Aplica la infraestructura:
   ```bash
   terraform apply
   ```
5. Sigue las instrucciones de salida para obtener las IPs/URLs de acceso.

### Notas

- Asegúrate de tener configuradas tus credenciales de AWS antes de aplicar Terraform.
- El despliegue creará VPC, subredes, gateway, security groups, etc.
- Asegurate de tener los secretos necesarios en AWS Secrets Manager, especialmente `jwt-secret` con los valores para `JWT_SECRET` y `NODE_APP_URL`.

---

## Scripts Útiles

- `scripts/build_push_express.sh`: Build y push de la imagen Express
- `scripts/build_push_fiber.sh`: Build y push de la imagen Fiber

---
