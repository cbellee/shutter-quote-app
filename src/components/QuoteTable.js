import React, { useState } from "react";
import { Link } from "react-router-dom";
import { Segment, Table, Button, Icon } from "semantic-ui-react";
import { GetQuotesFromLocalStorage, RemoveQuoteFromLocalStorage } from "../helpers/Helpers";
import seedData from '../helpers/SeedData';

export default function QuoteTable() {
    var quotes = GetQuotesFromLocalStorage('quotes', seedData);
    const [selectedId, setSelectedId] = useState();

    if (!quotes || quotes === null || quotes === '') {
        return (
            <h2>No Data Found...</h2>
        )
    }

    const onRemoveQuote = () => {
        RemoveQuoteFromLocalStorage('quotes', selectedId)
        setSelectedId(null); // Clear selection
    };

    const onEditQuote = () => {
        setSelectedId(null); // Clear selection
    };

    const rows = quotes.map(quote => (
        <Table.Row
            key={quote.id}
            onClick={() => setSelectedId(quote.id)}
            active={quote.id === selectedId}>
            <Table.Cell>{quote.id}</Table.Cell>
            <Table.Cell>{quote.dateCreated}</Table.Cell>
            <Table.Cell>{quote.firstName} {quote.lastName}</Table.Cell>
            <Table.Cell>{quote.email}</Table.Cell>
            <Table.Cell>{quote.windows.length}</Table.Cell>
            <Table.Cell>${quote.total}</Table.Cell>
        </Table.Row>
    ));

    return (
        <Segment>
            <Table celled striped selectable sortable>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>Id</Table.HeaderCell>
                        <Table.HeaderCell>Date Created</Table.HeaderCell>
                        <Table.HeaderCell>Name</Table.HeaderCell>
                        <Table.HeaderCell>Email</Table.HeaderCell>
                        <Table.HeaderCell>Windows</Table.HeaderCell>
                        <Table.HeaderCell>Total</Table.HeaderCell>
                    </Table.Row>
                </Table.Header>
                <Table.Body>{rows}</Table.Body>
                <Table.Footer fullWidth>
                    <Table.Row>
                        <Table.HeaderCell />
                        <Table.HeaderCell colSpan="6">
                            <Link to={`/quote/${selectedId}`}>
                                <Button
                                    floated="left"
                                    icon
                                    labelPosition="left"
                                    color="green"
                                    size="small"
                                    disabled={!selectedId}
                                    onClick={onEditQuote}>
                                    <Icon name="edit" /> Edit Quote
                            </Button>
                            </Link>
                            <Button
                                floated="left"
                                icon
                                labelPosition="left"
                                color="red"
                                size="small"
                                disabled={!selectedId}
                                onClick={onRemoveQuote}>
                                <Icon name="trash" /> Remove Quote
                            </Button>
                        </Table.HeaderCell>
                    </Table.Row>
                </Table.Footer>
            </Table>
        </Segment>
    );
}
