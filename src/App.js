import React, { useState } from 'react';
import './App.css';
import QuoteForm from './components/QuoteForm';
import Header from './components/Header'
import Footer from './components/Footer'
import MainMenu from './components/MainMenu'
import PostCodeContext from './components/PostCodeContext';
import PostCodeData from './helpers/PostCodeData';
import { GetQuotesFromLocalStorage } from './helpers/Helpers';
import { Message } from 'semantic-ui-react';

function App() {
  const postCodeDataHook = useState(PostCodeData);

  return (
    <html>
      <PostCodeContext.Provider value={postCodeDataHook}>
        <div className="App">
          <Header />
          <MainMenu />
          <Footer />
        </div>
      </PostCodeContext.Provider>
    </html>
  );
}

export default App;
