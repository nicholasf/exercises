import express, {} from 'express';
import * as jwt from 'jsonwebtoken';
let service = express();
const hello = (req, res, next) => {
    let token = jwt.sign({ foo: "bar" }, "shhh");
    console.log("token: ", token);
    res.setHeader("Authorization", "Bearer " + token.toString());
    res.json({ message: "Hello!" });
};
service.get("/hello", hello);
service.listen(3000, () => {
    console.log("Listening on 3000");
});
//# sourceMappingURL=index.js.map