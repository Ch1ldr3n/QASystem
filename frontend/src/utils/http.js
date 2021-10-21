import axios from "axios";

const instance = axios.create({
	baseURL:"v1/",
	timeout: 5000,
})
/*
export function Get(translate) {
    axios.get()
    .then(translate)
	.catch(function (error) {
    console.log(error);
    });
}
*/

export function postsignin(username,password,token,translate) {
    instance.post("user/login",{
        username: username,
        password: password
    })
    .then(translate)
    .catch(function (error) {
       console.log(error);
       alert("登录失败")
    });
}

export function postregister(username,password,translate) {
    instance.post("user/register",{
        username: username,
        password: password
    })
    .then(translate)
    .catch(function (error) {
       console.log(error);
       alert("注册失败")
    });
}