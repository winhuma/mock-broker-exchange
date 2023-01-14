import Router from "next/router";
import React from "react";
import {useState, useEffect} from "react"
import { Navbar, Button, Card, Grid, Text, Radio, Dropdown } from "@nextui-org/react";
import { Modal, Input, Row } from "@nextui-org/react";
import axios from 'axios';

function Exchange() {
    if (typeof window !== 'undefined') {
        const myToken = localStorage.getItem("mytoken")
        if (myToken == null) {
            Router.replace("/")
        }
    }
    function logoutClick() {
        Router.replace('/');
    }
    const [currency, setCurrency] = useState([]);
    const [mebalance, setMeBalance] = useState([]);
    const [orderCurrencyTarget, setCurrencyValue] = useState([]);
    const [visible, setVisible] = React.useState(false);
    const [visibleSuccess, setVisibleSuccess] = React.useState(false);
    const [actionType, setActionType] = useState('');
    const [amountAvailable, setAmountAvailable] = useState('');
    const [amountReceive, setAmountReceive] = useState('');
    const [selected, setSelected] = React.useState(new Set(["select your balance"]));
    const [rate1unit, setRate1Unit] = useState('');
    const [amountOrder, setAmountOrder] = useState(0);
    const [currencyIDPaySelect, setCurrencyIDPaySelect] = useState(0);
    const [currencyIDTargetSelect, setCurrencyIDTargetSelect] = useState(0);
    const [myUsername, setMyUsername] = useState('');

    const selectedMeValue = React.useMemo(
        () => Array.from(selected).join(", ").replaceAll("_", " "),
        [selected]
      );
    
    const amountOrderChange = (event: { target: { value: string; }; }) => {
        setAmountOrder(parseFloat(event.target.value))
        var balanceData = parseFloat(event.target.value)
        var targetRate: number
        var mePayRate: number
        var calResult: number
        if (setActionType == "BUY") {
            for (let j = 0; j < currency.length; j++) {
                if (currency[j].name == orderCurrencyTarget) {
                    setCurrencyIDPaySelect(currency[j].id)
                    targetRate = parseFloat(currency[j].value_usd)
                }
                if (currency[j].name == selectedMeValue) {
                    setCurrencyIDTargetSelect(currency[j].id)
                    mePayRate = parseFloat(currency[j].value_usd)
                }
            }
        } else {
            for (let j = 0; j < currency.length; j++) {
                if (currency[j].name == orderCurrencyTarget) {
                    setCurrencyIDTargetSelect(currency[j].id)
                    targetRate = parseFloat(currency[j].value_usd)
                }
                if (currency[j].name == selectedMeValue) {
                    setCurrencyIDPaySelect(currency[j].id)
                    mePayRate = parseFloat(currency[j].value_usd)
                }
            }
        }
        
        
        if (actionType == "BUY") {
            calResult = mePayRate / targetRate
            const result = balanceData * calResult
            setAmountReceive(result.toString() + " " + orderCurrencyTarget)
        } else {
            calResult = mePayRate / targetRate
            const result = balanceData / calResult
            setAmountReceive(result.toString() + " " + selectedMeValue)
        }
        
    }
    const actionChange = (event:string) => {
        setActionType(event)
        var rateTarget: number
        var rateMyMustPay: number
        for (let i = 0; i < currency.length; i++) {
            if (currency[i].name == orderCurrencyTarget) {
                rateTarget = parseFloat(currency[i].value_usd)
            }
            if (currency[i].name == selectedMeValue) {
                rateMyMustPay = parseFloat(currency[i].value_usd)
            }
        }
        var rate2show = rateMyMustPay / rateTarget
        setRate1Unit(`1 ${orderCurrencyTarget} = ${rate2show.toString()} ${selectedMeValue}`)
        let found:boolean
        if (event == "SALE") {
            for (let i = 0; i < mebalance.length; i++) {
                if (mebalance[i].currency_name === orderCurrencyTarget) {
                    setAmountAvailable(mebalance[i].balance + " " + orderCurrencyTarget);
                    found=true
                }
                
            }
        } else {
            
            for (let i = 0; i < mebalance.length; i++) {
                if (selectedMeValue == orderCurrencyTarget) {
                    setAmountAvailable("this sane currency");
                    found=true
                }
                if (mebalance[i].currency_name === selectedMeValue) {
                    setAmountAvailable(mebalance[i].balance + " " + selectedMeValue);
                    found=true
                }   
            }
        }
        if (!found) {
            setAmountAvailable("not found");
        }
    }
    useEffect(() => {
        async function fetchData() {
            try {
                setMyUsername(JSON.parse(localStorage.getItem("myusername") || '{}'))
            } catch (error) {
                console.error(error);
            }
        }
        fetchData();
    }, []);
    useEffect(() => {
        async function fetchData() {
            try {
                const response = await axios.get(process.env.NEXT_PUBLIC_BACKEND_API + "/currencies");
                setCurrency(response.data.data);
            } catch (error) {
                console.error(error);
            }
        }
        fetchData();
    }, []);
    useEffect(() => {
        async function fetchData() {
            try {
                const myToken = localStorage.getItem("mytoken") || ''
                const response = await axios.get(process.env.NEXT_PUBLIC_BACKEND_API + "/balance/me", {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${myToken}`
                    },
                    responseType: 'json'
                });
                setMeBalance(response.data.data);
            } catch (error) {
                console.error("asdsadasda", error);
            }
        }
        fetchData();
    }, []);

    function currencyClick(cur:any) {
        setCurrencyValue(cur)
        setVisible(true);
    }
    const closeHandler = () => {
        setVisible(false);
    };

    const closeSuccess = () => {
        setVisibleSuccess(false)
    };
    const createOrder = () => {
        const myToken = localStorage.getItem("mytoken") || ''
        axios({
            method: 'post',
            url: process.env.NEXT_PUBLIC_BACKEND_API + "/orders",
            data: {
                my_currency_id: currencyIDPaySelect,
                target_currency_id: currencyIDTargetSelect,
                my_currency_value: amountOrder,
                action: actionType, 
            },
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${myToken}`
            },
            responseType: 'json'
        }).then((response: any) => {
            if (response.status == 200) {
                axios.get(process.env.NEXT_PUBLIC_BACKEND_API + "/balance/me", {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${myToken}`
                    },
                    responseType: 'json'
                }).then((resNewBalance: any) => {
                    setMeBalance(resNewBalance.data.data);
                });   
            }
        });
        setVisible(false);
        setVisibleSuccess(true)
    }
    return (
        <div>
            <Navbar isBordered variant="sticky">
                <Navbar.Brand>
                </Navbar.Brand>
                <Navbar.Content>
                    <Navbar.Link color="inherit">
                    <h3>User: {myUsername}</h3>
                    </Navbar.Link>
                    <Navbar.Item>
                    <Button shadow auto color="error" onPress={logoutClick}>
                        Logout
                    </Button>
                    </Navbar.Item>
                </Navbar.Content>
            </Navbar>
            <br/>
            <Card style={{background: "red", width:"80%", margin: "auto"}}>
                <br/>
                <center><h1>My Balance</h1></center>
                <Grid.Container style={{margin: "auto"}} gap={2} justify="flex-start">
                    {mebalance.map((item, index) => (
                        <Grid style={{margin: "auto"}} xs={6} sm={3} key={index}>
                        <Card css={{ p: "$6", mw: "300px" }}>
                            <Card.Body css={{ p: 0, ta: "center" }}>
                            <h2>{item.currency_name}</h2>
                            <h2>{item.balance}</h2>
                            </Card.Body>
                        </Card>
                        </Grid>
                    ))}
                </Grid.Container>
            </Card>
            <br/>
            <div style={{margin: "auto"}}>
                <center>
                    {currency.map((item, index) => (
                            <Button
                            key={index}
                            style={{width:"50%"}}
                            shadow
                            color="gradient"
                            size="xl"
                            value={item.name}
                            onClick={() => currencyClick(item.name)}>
                                {item.name} : {item.value_usd} / USD</Button>
                    ))}
                </center>
            </div>
            <div>
      <Modal
        closeButton
        preventClose
        aria-labelledby="modal-title"
        open={visible}
        onClose={closeHandler}
      >
        <Modal.Header>
          <Text id="modal-title" size={18}>
            Order {orderCurrencyTarget} {actionType}
          </Text>
        </Modal.Header>
        <Modal.Body>
          <Input
            clearable
            bordered
            fullWidth
            color="primary"
            size="lg"
            placeholder="Enter amount"
            onChange={amountOrderChange}
          />
          <Dropdown>
            <Dropdown.Button flat color="secondary" css={{ tt: "capitalize" }}>
                {selectedMeValue}
            </Dropdown.Button>
            <Dropdown.Menu
                aria-label="Single selection actions"
                color="secondary"
                disallowEmptySelection
                selectionMode="single"
                onSelectionChange={setSelected}>
                {currency.map((item, _) => (
                    <Dropdown.Item key={item.name}>
                        {item.name}
                    </Dropdown.Item>
                ))}
            </Dropdown.Menu>
            </Dropdown>
          <div style={{textAlign: "right"}}>available : {amountAvailable}</div>
          <div style={{textAlign: "right"}}>be receive : {amountReceive}</div>
          <div style={{textAlign: "right"}}>Rate : {rate1unit}</div>
          <Row justify="space-between">
            <Radio.Group orientation="horizontal" label="Action" value={actionType} onChange={actionChange}>
                <Grid>
                    <Radio color="success" labelColor="success" value="BUY">
                        BUY
                    </Radio>
                </Grid>
                <Grid>
                    <Radio color="error" labelColor="error" value="SALE">
                        SALE
                    </Radio>
                </Grid>
            </Radio.Group>
          </Row>
        </Modal.Body>
        <Modal.Footer>
          <Button auto flat color="error" onPress={closeHandler}>
            Close
          </Button>
          <Button auto onPress={createOrder}>
            Create Order
          </Button>
        </Modal.Footer>
      </Modal>



      <Modal
        closeButton
        preventClose
        aria-labelledby="modal-title"
        open={visibleSuccess}
        onClose={closeSuccess}
      >
        <Modal.Body>
            SUCCESS
        </Modal.Body>
        <Modal.Footer>
          <Button auto onPress={closeSuccess}>
            OK
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
        </div>
        
    )
}
export default Exchange