import express, { type Request, type Response, type NextFunction } from 'express';
import fs from 'node:fs';
import jwt from 'jsonwebtoken';

let service = express();

const helloBase64 = (req: Request, res: Response, next: NextFunction) => {
    let token = jwt.sign({ foo: "bar"}, "shhh")
    console.log("token: ", token)
    res.setHeader("Authorization", "Bearer " + token)
    res.json({ message: "Hello!" })
}

const privateKey = fs.readFileSync('./../private-key.pem');
// console.log('private key: ', privateKey.toString())

const helloWithEncryption = (req: Request, res: Response, next: NextFunction) => {
    let token = jwt.sign({ foo: "bar" }, privateKey, { algorithm: "RS256" })
    res.setHeader("Authorization", "Bearer " + token)
    res.json({ message: "Hello2!" })
}

const helloWithCookie = (req: Request, res: Response, next: NextFunction) => {
    res.cookie("username", "fred", { maxAge: 360})
    res.cookie("postcode", "3000")
    res.json({ message: "Hello2!" })
}

service.get("/hello", helloBase64)
service.get("/hello2", helloWithEncryption)
service.get("/hello3", helloWithCookie)

service.listen(3000, ()=> {
    console.log("Listening on 3000")
})