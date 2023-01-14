import { Inter } from '@next/font/google'
import { Button, Input, Card } from "@nextui-org/react";
import Router from "next/router";

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
    var userNumber: string
    const UserIDChange = (event: { target: { value: string; }; }) => {
        userNumber = event.target.value
    }
    const MyLogin = () => {

        const axios = require('axios');
        if (userNumber != "") {
            axios({
                method: 'post',
                url: process.env.NEXT_PUBLIC_BACKEND_API + "/login",
                data: {
                    username: userNumber,
                },
                responseType: 'json'
            }).then((response: any) => {
                console.log(response.data.data, typeof(response.data.data.user_balance), userNumber, response.data.data.user_balance)
                localStorage.setItem("mytoken", response.data.data.token)
                localStorage.setItem("myusername", userNumber)
                if (response.data.data.user_balance == null) {
                    Router.replace("/newuser/starter")
                } else {
                    Router.replace("/exchange/exchange")    
                }
            });
        }
    }
    return (
        <div>
            <br/><br/><br/><br/>
            <center>
                <Card css={{ p: "$20", mw: "600px" }}>
                    <h1>Wanna be Exchange!!!</h1>
                    <br/><br/><br/>
                    <Input
                        clearable
                        underlined
                        status="default" 
                        color="default"
                        size="lg"
                        onChange={UserIDChange}
                        labelPlaceholder="Enter Your ID here..." />
                    <br/><br/>
                    <Button
                        shadow color="success" auto
                        onPress={MyLogin}
                    >Enter</Button>
                </Card>
            </center>
        </div>
    )
}