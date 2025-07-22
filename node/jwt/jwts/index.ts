import express, { type Request, type Response, type NextFunction } from 'express';
import jwt from 'jsonwebtoken';

let service = express();

const hello = (req: Request, res: Response, next: NextFunction) => {
    let token = jwt.sign({ foo: "bar"}, "shhh")
    console.log("token: ", token)
    res.setHeader("Authorization", "Bearer " + token.toString())
    res.json({ message: "Hello!" })
}

const hello2 = (req: Request, res: Response, next: NextFunction) => {
    res.cookie("username", "fred", { maxAge: 360})
    res.cookie("postcode", "3000")
    res.json({ message: "Hello2!" })
}

service.get("/hello", hello)
service.get("/hello2", hello2)

service.listen(3000, ()=> {
    console.log("Listening on 3000")
})