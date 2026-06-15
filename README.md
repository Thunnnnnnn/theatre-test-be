# System Architecture Diagram

```text
                    ┌───────────────┐
                    │   Frontend    │
                    │  Nuxt 3 / Vue │
                    └───────┬───────┘
                            │ HTTP
                            ▼
                    ┌───────────────┐
                    │    Go API     │
                    │ Gin Framework │
                    └───────┬───────┘
                            │
         ┌──────────────────┼──────────────────┐
         │                  │                  │
         ▼                  ▼                  ▼
 ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
 │  MongoDB    │   │    Redis    │   │    Kafka    │
 │ Primary DB  │   │    Lock     │   │ Event Queue │
 └─────────────┘   └─────────────┘   └──────┬──────┘
                                            │
                                            ▼
                                    ┌─────────────┐
                                    │   Worker    │
                                    │ Background  │
                                    │ Processing  │
                                    └──────┬──────┘
                                           │
                                           ▼
                                    ┌─────────────┐
                                    │ WebSocket   │
                                    │ Realtime    │
                                    └─────────────┘
```

---

# Tech Stack Overview

### Frontend

Frontend git url: https://github.com/Thunnnnnnn/theatre-test-fe.git

* Nuxt 3
* Vue 3
* Pinia
* Bootstrap 5
* SweetAlert2
* BootstrapVue

### Backend

* Go
* Gin Framework
* Gorilla WebSocket
* JWT Authentication

### Database

* MongoDB

### Infrastructure

* Docker
* Docker Compose

### Messaging & Realtime

* Apache Kafka
* WebSocket

### Distributed Systems

* Redis
* Distributed Lock

---

# ขั้นตอนการจอง

1. เข้าสู่ระบบ โดยใช้ gmail ในการเข้าสู่ระบบ
2. ไปที่ http://localhost:3000/theatres เพื่อเลือกโรงหนังที่ต้องการจอง
3. เมื่อเลือกโรงหนังแล้ว จะแสดงข้อมูลเก้าอี้ที่ว่าง ที่ถูกจองไว้ก่อนชำระเงิน และที่ชำระเงินสำเร็จ โดยจะแยกตามสีแต่ละประเภท
4. เมื่อเลือกที่นั่งแล้ว จะมี Modal ขึ้นมายืนยันการจองก่อนชำระเงิน ให้กดตกลงเพื่อจองไว้ก่อนชำระเงิน โดยจะต้องชำระภายใน 5 นาทีหลังกดจอง
5. ให้ไปที่ http://localhost:3000/my-booked เพื่อไปชำระเงิน โดยเก้าอี้ที่จองไว้แต่ยังไม่ได้ชำระเงิน จะมีปุ่มให้กดเพื่อยืนยันการชำระเงิน
6. หลังจากกดปุ่มชำระเงิน จะมี Modal เป็น Mock modal เพื่อให้ยืนยันชำระเงิน
7. หลังจากชำระเงินสำเร็จก็เป็นอันเสร็จสิ้น

# Redis Lock Strategy

Redis lock ของระบบนี้ จะล็อคที่นั่งที่ถูกจองโดยผู้ใช้คนแรก และหลังจากนั้นที่นั่งนั้นจะโดนล็อคเป็นเวลา 10 วิ เพื่อกันการจองซ้อนกัน

# Message Queue Usage

Message queue ในระบบนี้นำไปใช้ในการส่ง Websocket หลังจากที่จองที่นั่งสำเร็จ และหลังจากหมดเวลาชำระเงิน เพื่อให้หน้าเว็บได้มีการอัพเดทเป็นปัจจุบันในส่วนของหน้าการจองที่นั่ง 

# วิธีรันระบบ
### สร้างไฟล์ .env
สร้างไฟล์ .env ใน root project แบบ .env.sample และกรอกข้อมูลของ env

### Start Services
```bash
docker compose up -d
```

### API Endpoint

```text
http://localhost:8080
```

### WebSocket Endpoint

```text
ws://localhost:8080/ws
```

### Stop Services

```bash
docker compose down
```

---

# Assumptions & Trade-offs

---

