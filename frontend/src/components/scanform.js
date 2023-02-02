import React, { useState } from 'react';

const ScanForm = () => {
  const [address, setAddress] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();

    const goURL = 'http://localhost:3200';
    const hasuraURL = 'http://localhost:8080/v1/graphql';
    const graphlresultaddr = document.getElementById("graphlresultaddr");
    const graphlresulttxcount = document.getElementById("graphlresulttxcount");
    graphlresultaddr.innerHTML = "Getting Data...";

    fetch(goURL, {
      method: 'POST',
      headers: {
        'Content-Type': 'multipart/form-data'
      },
      body: 'address=' + address 
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
        graphlresulttxcount.innerHTML = data
      });

    // body: 'query { contracts (where: {address: {_eq: "' + address + '"}}) { address transactions } operations (where: {from: {_eq: "' + address + '"}, _or: {to: {_eq: "' + address + '"}}}) {from hash to value timestamp}}'

    const query = `query { contracts (where: {address: {_eq: "` + address + `"}}) { address transactions }}`;
    console.log(query)

    fetch(hasuraURL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'X-Hasura-Admin-Secret': 'mysecretkey',
        'Access-Control-Allow-Origin': '*'
      },
      body: JSON.stringify({
        query
      })
    })
      .then(response => response.json())
      .then(data => {
        console.log(data.data["contracts"][0]);
        console.log(data.data["contracts"][0].address);
        console.log(data.data["contracts"][0].transactions);
        graphlresultaddr.innerHTML = data.data["contracts"][0].address;
        graphlresulttxcount.innerHTML = data.data["contracts"][0].transactions;
      });

  }
  return (
    <div>
      <img src="https://ethplorer.io/wallet/env/ethplorer/img/ethplorer-blue.png" width="25%" className="App-logo" alt="logo" />
      <form onSubmit={handleSubmit}>
        <p>
          <input
            type="text"
            value={address}
            onChange={e => setAddress(e.target.value)}
          />
        </p>
        <p>
          <button type="submit">Scan</button>
        </p>
      </form>
      <p id="graphlresultaddr">
      </p>
      <p id="graphlresulttxcount">
      </p>
    </div>
  );
};

export default ScanForm;
