import React from "react";
import { Route, NavLink, HashRouter } from "react-router-dom";
import { Menu, MenuItem, Container } from "semantic-ui-react";
import Home from './Home';
import Contact from './Contact';
import QuoteForm from './QuoteForm';
import QuoteTable from './QuoteTable';
import QuoteDetailForm from './QuoteDetailForm';

function MainMenu({ quotes, customers }) {

    return (
        <HashRouter>
            <div>
                <Menu>
                    <Container>
                        <MenuItem><NavLink exact to="/">Quotes</NavLink></MenuItem>
                        <MenuItem><NavLink to="/newquote">Add Quote</NavLink></MenuItem>
                        <MenuItem><NavLink to="/contact">Contact</NavLink></MenuItem>
                    </Container>
                </Menu>
                <Route exact path="/" component={QuoteTable} />
                <Route path="/quote/:quoteId" component={QuoteDetailForm} />
                <Route path="/newquote" component={QuoteForm} />
                <Route path="/contact" component={Contact} />
            </div>
        </HashRouter>
    );
}

export default MainMenu;