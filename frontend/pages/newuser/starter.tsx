import Link from 'next/link'
import Router from "next/router";
import { Button, Input, Card, useInput, Dropdown } from "@nextui-org/react";
import { text } from 'stream/consumers';
import React from "react";
import {useState, useEffect} from "react"
import axios from 'axios';

function Starter() {
    if (typeof window !== 'undefined') {
        const myToken = localStorage.getItem("mytoken")
        if (myToken == null) {
            Router.replace("/")
        }
    }

    const [mybalance, setMybalance] = useState('');
    const { value, reset, bindings } = useInput("");
    const [options, setOptions] = useState([]);
    const validateBalance = (value: string) => {
        return value.match(/^\d*\.?\d*$/i);
    };
    const helper = React.useMemo(() => {
        if (!value)
            return {
            text: "",
            color: "",
            };
        const isValid = validateBalance(value);
        setMybalance(value)
        return {
            text: isValid ? "Correct format" : "Enter a valid number format",
            color: isValid ? "success" : "error",
        };
    }, [value]);

    useEffect(() => {
        async function fetchData() {
            try {
                const response = await axios.get(process.env.NEXT_PUBLIC_BACKEND_API + "/currencies");
                setOptions(response.data.data);
            } catch (error) {
                console.error(error);
            }
        }
        fetchData();
    }, []);

    var currencyID: number
    const CurrencyChange = (event: { target: { value: number; }; }) => {
        currencyID = event.target.value
    }

    const SendStarter = () => {

        const myToken = localStorage.getItem("mytoken") || ''
        if (currencyID != 0) {
            axios({
                method: 'post',
                url: process.env.NEXT_PUBLIC_BACKEND_API + "/balance/starter",
                data: {
                    balance_start: mybalance,
                    currency_id: currencyID,
                },
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${myToken}`
                },
                responseType: 'json'
            }).then((response: any) => {
                if (response.status == 200) {
                    Router.replace("/exchange/exchange")
                }
            });
        }
    }
    return (
        <div>
            <Link href="/"><h1>Back to index</h1></Link>
            <h1>Starter</h1>
            <center>
                <Card css={{ p: "$15", mw: "400px" }}>
                    <h1>You are new user</h1>
                    <br/>
                    <Input
                        {...bindings}
                        clearable
                        shadow={false}
                        onClearClick={reset}
                        status={helper.color}
                        color={helper.color}
                        helperColor={helper.color}
                        helperText={helper.text}
                        type="text"
                        placeholder="Enter starter balance"
                        />
                    <br/><br/>
                    
                    <select onChange={CurrencyChange}>
                        {options.map((option, index) => (
                            <option key={index} value={option.id}>{option.name}</option>
                        ))}
                    </select>
                    <Button
                        onClick={SendStarter}
                        shadow color="success" auto
                    >Enter</Button>
                </Card>
            </center>
            
        </div>
    )
}
export default Starter