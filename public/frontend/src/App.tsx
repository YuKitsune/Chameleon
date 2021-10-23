import React from 'react';
import {BrowserRouter, Switch, Route} from 'react-router-dom';
import Navbar from './components/header/Navbar';
import Aliases from './components/pages/aliases';
import QuarantinedMail from './components/pages/quarantinedMail';
import Account from './components/pages/account';
import Settings from './components/pages/settings';
import './App.css';

function App() {
  return (
    <div className='flex flex-col h-screen overflow-hidden bg-gray-100'>
      <BrowserRouter>
        <Navbar />

        <div className='content px-6 lg:px-24'>
          <main className='child'>

            <Switch>
              <Route exact path="/aliases">
                <Aliases />
              </Route>
              <Route exact path="/quarantined">
                <QuarantinedMail />
              </Route>
              <Route exact path="/account">
                <Account />
              </Route>
              <Route exact path="/settings">
                <Settings />
              </Route>
            </Switch>

          </main>

          <footer className='footer'>
            <p>visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to learn SvelteKit</p>
          </footer>

        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
