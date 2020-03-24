import React, { useState, useContext } from 'react';
import { useForm } from 'react-hook-form';
import { Redirect } from 'react-router-dom';
import WindowInputs from './WindowInputs';
import { AddQuoteToLocalStorage, GetQuotesFromLocalStorage, GetPostCodeFromSuburb, GetSuburbList } from '../helpers/Helpers';
import { Segment, Icon, Form, Input, Button, Grid, Card, GridRow, GridColumn } from "semantic-ui-react";
import seedData from '../helpers/SeedData';
import PostCodeContext from './PostCodeContext';

const dateCreated = new Date().toISOString();

export default function QuoteForm() {
    const postCodes = useContext(PostCodeContext)[0];
    const { register, handleSubmit, watch, errors } = useForm()
    var stateName = "NSW"
    var suburbs = GetSuburbList(postCodes, stateName);

    const [quoteState, setQuoteState] = useState({
        id: '',
        dateCreated: dateCreated,
        total: 0,
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        street: '',
        suburb: '',
        city: '',
        postCode: '',
        windows: []
    });

    const handleQuoteChange = (e) => setQuoteState({
        ...quoteState,
        [e.target.name]: e.target.value,
    });

    const handleSuburbChange = (e) => setQuoteState({
        ...quoteState,
        postCode: GetPostCodeFromSuburb(postCodes, e.target.value)
    });

    const blankWindow = { name: '', width: '2400', height: '1200', panel: '2' };
    const [windowState, setWindowState] = useState([
        { ...blankWindow },
    ]);

    const [redirectToQuotes, setRedirectToQuotes] = useState(false);

    const addWindow = () => {
        setWindowState([...windowState, { ...blankWindow }]);
    };

    const handleWindowChange = (e) => {
        const updatedWindows = [...windowState];
        console.log(JSON.stringify(updatedWindows));
        updatedWindows[e.target.dataset.idx][e.target.className] = e.target.value;
        setWindowState(updatedWindows);
    };

    const handleWindowRemove = (e) => {
        const updatedWindows = [...windowState];
        console.log("updated windows: " + JSON.stringify(updatedWindows));
        console.log("deleted window index: " + e.target.dataset.idx);
        updatedWindows.splice(e.target.dataset.idx, 1);
        setWindowState(updatedWindows);
        console.log("updated windows: " + JSON.stringify(updatedWindows));
    }

    const onSubmit = data => {
        quoteState.windows = windowState;
        var totalCost = 0;
        quoteState.windows.map((window) => {
            if (parseInt(window.price) != '') {
                totalCost = totalCost + parseInt(window.price);
            }
        });
        quoteState.total = totalCost;
        AddQuoteToLocalStorage('quotes', quoteState, seedData);
        setRedirectToQuotes(true);
        console.log(data);
    }

    return (
        <Form>
            {redirectToQuotes ? <Redirect to="/" /> : null}
            <Segment>
                <Grid columns={3} padded="vertically">
                    <Grid.Column>
                        <Form.Field>
                            <label>First Name</label>
                            <input
                                label="First Name"
                                type="text"
                                name="firstName"
                                id="firstName"
                                ref={register({ required: true })}
                                value={quoteState.firstName}
                                onChange={handleQuoteChange}
                            />
                            {errors.firstName && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Last Name</label>
                            <input
                                label="Last Name"
                                type="text"
                                name="lastName"
                                id="lastName"
                                ref={register({ required: true, minLength: 3, maxLength: 40 })}
                                value={quoteState.lastName}
                                onChange={handleQuoteChange}
                            />
                            {errors.lastName && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Email</label>
                            <input
                                label="Email"
                                type="text"
                                name="email"
                                id="email"
                                ref={register({ required: true, pattern: /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/ })}
                                value={quoteState.email}
                                onChange={handleQuoteChange}
                            />
                            {errors.email && <span>valid email address is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Phone</label>
                            <input
                                label="Phone"
                                type="text"
                                name="phone"
                                id="phone"
                                ref={register({ required: true, minLength: 10, maxLength: 14 })}
                                value={quoteState.phone}
                                onChange={handleQuoteChange}
                            />
                            {errors.phone && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Street</label>
                            <input
                                label="Street"
                                type="text"
                                name="street"
                                id="street"
                                ref={register({ required: true, minLength: 3, maxLength: 40 })}
                                value={quoteState.street}
                                onChange={handleQuoteChange}
                            />
                            {errors.street && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Suburb</label>
                            <select
                                label="Suburb"
                                name="suburb"
                                id="suburb"
                                value={quoteState.suburb}
                                onChange={handleSuburbChange}>
                                {suburbs.map((suburb) => <option key={suburb.text} value={suburb.value}>{suburb.value}</option>)}
                            </select>
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>PostCode</label>
                            <input
                                label="Postcode"
                                type="text"
                                name="postCode"
                                id="postCode"
                                value={quoteState.postCode}
                                onChange={handleQuoteChange}
                            />
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Total</label>
                            <input
                                label="Total"
                                type="text"
                                name="total"
                                id="total"
                                value={quoteState.total}
                                onChange={handleQuoteChange}
                            />
                        </Form.Field>
                    </Grid.Column>
                </Grid>
            </Segment>
            <Grid>
                <Form.Field>
                    <Button
                        floated="left"
                        icon
                        labelPosition="left"
                        color="blue"
                        size="small"
                        onClick={addWindow}>
                        <Icon name="edit" /> Add Window
                        </Button>
                </Form.Field>
                <Form.Field>
                    <Button
                        floated="left"
                        icon
                        labelPosition="left"
                        color="green"
                        size="small"
                        onClick={handleSubmit(onSubmit)}>
                        <Icon name="edit" /> Save Quote
                        </Button>
                </Form.Field>
            </Grid>
            <Grid columns={1}>
                <Grid.Column>
                    {
                        windowState.map((val, idx) => (
                            <WindowInputs
                                key={`window-${idx}`}
                                idx={idx}
                                windowState={windowState}
                                handleWindowChange={handleWindowChange}
                                handleWindowRemove={handleWindowRemove}
                            />
                        ))
                    }
                </Grid.Column>
            </Grid>
        </Form >
    );
};