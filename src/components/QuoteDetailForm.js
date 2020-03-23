import React, { useState, useContext } from 'react';
import { Redirect } from 'react-router-dom';
import WindowInputs from './WindowInputs';
import { UpdateQuoteToLocalStorage, GetQuotesFromLocalStorage, GetPostCodeFromSuburb, GetSuburbList } from '../helpers/Helpers';
import { Segment, Form, Icon, Input, Button, Grid, Card, GridRow, GridColumn } from "semantic-ui-react";
import seedData from '../helpers/SeedData';
import PostCodeContext from './PostCodeContext';

const dateCreated = new Date().toISOString();

const QuoteDetailForm = ({ match }) => {
    const {
        params: { quoteId },
    } = match;

    let quotes = GetQuotesFromLocalStorage('quotes', seedData);
    let quote = quotes.filter((quote) => quote.id == quoteId)[0];

    const postCodes = useContext(PostCodeContext)[0];
    var stateName = "NSW"
    var suburbs = GetSuburbList(postCodes, stateName);

    const [quoteState, setQuoteState] = useState(
        quote,
    );

    const [windowState, setWindowState] = useState(
        quote.windows,
    );

    const handleQuoteChange = (e) => setQuoteState({
        ...quoteState,
        [e.target.name]: e.target.value,
    });

    const handleSuburbChange = (e) => setQuoteState({
        ...quoteState,
        postCode: GetPostCodeFromSuburb(postCodes, e.target.value),
        [e.target.name]: e.target.value,
    });

    const handleTotalChange = (e) => {
        const updateWindows = [...windowState];
        quoteState.windows = windowState;
        console.log("totalCostQuote: " + JSON.stringify(quoteState));
        var totalCost = 0;
        quoteState.windows.map((window) => {
            totalCost = totalCost + parseInt(window.price);
        });
        console.log("totalCost: " + totalCost);
        quoteState.total = totalCost;
        setQuoteState(quoteState);
    }

    console.log("quoteId: " + quoteId);
    console.log("quote state: " + JSON.stringify(quoteState.quote));
    console.log("window state: " + JSON.stringify(windowState));

    const blankWindow = { name: '', width: '2400', height: '1200', panel: '2' };
    const addWindow = () => {
        setWindowState([...windowState, { ...blankWindow }]);
    };

    const [redirectToQuotes, setRedirectToQuotes] = useState(false);

    const handleWindowChange = (e) => {
        const updatedWindows = [...windowState];
        console.log(JSON.stringify(updatedWindows));
        updatedWindows[e.target.dataset.idx][e.target.className] = e.target.value;
        setWindowState(updatedWindows);
    };

    const handleWindowRemove = (e) => {
        const updatedWindows = [...windowState];
        console.log(JSON.stringify(updatedWindows));
        updatedWindows.splice(e.target.dataset.idx, 1);
        setWindowState(updatedWindows);
    }

    const onSubmit = () => {
        quoteState.windows = windowState;
        var totalCost = 0;
        quoteState.windows.map((window) => {
            totalCost = totalCost + parseInt(window.price);
        });
        quoteState.total = totalCost;
        UpdateQuoteToLocalStorage('quotes', quoteState, seedData);
        setRedirectToQuotes(true);
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
                                value={quoteState.firstName}
                                onChange={handleQuoteChange}
                            />
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
                                value={quoteState.lastName}
                                onChange={handleQuoteChange}
                            />
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
                                value={quoteState.email}
                                onChange={handleQuoteChange}
                            />
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
                                value={quoteState.phone}
                                onChange={handleQuoteChange}
                            />
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
                                value={quoteState.street}
                                onChange={handleQuoteChange}
                            />
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
                                onChange={handleTotalChange}
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
                        onClick={onSubmit}>
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

export default QuoteDetailForm;