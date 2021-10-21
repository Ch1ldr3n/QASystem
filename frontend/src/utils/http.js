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
	document.cookie = "token=" + token +";expires=Fri, 7-July-2023 01:01:01 GMT"
    instance.post("user/login",{
        username: username,
        password: password
    })
    .then(translate)
    .catch(function (error) {
    console.log(error);
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
    });
}