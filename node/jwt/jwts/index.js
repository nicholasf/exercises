"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express_1 = require("express");
var jsonwebtoken_1 = require("jsonwebtoken");
var sign = jsonwebtoken_1.default.sign;
var service = (0, express_1.default)();
var hello = function (req, res, next) {
    var token = sign({ foo: "bar" }, "shhh");
    console.log("token: ", token);
    res.setHeader("Authorization", "Bearer " + token.toString());
};
service.get("/hello", hello);
service.listen(3000, function () {
    console.log("Listening on 3000");
});
