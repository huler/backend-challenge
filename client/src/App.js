import "./App.css";

import React, { useState } from "react";
import Page1 from "./page1";
import Page2 from "./page2";
import Page2a from "./page2a";
import Page3 from "./page3";
import Choice from "./choice";
import Results from "./results";

function App() {
  let [appState, setAppState] = useState(null);
  let [currentPage, setCurrentPage] = useState(0);

  let [userDetails, setUserDetails] = useState(null);
  let [matrixData, setMatrixData] = useState(null);

  const handleNextPage = () => {
    currentPage++;
    setCurrentPage(currentPage);
  };
  const handlePreviousPage = () => {
    currentPage--;
    setCurrentPage(currentPage);
  };
  const handleSubmit = () => {
    handleNextPage();
  };

  if (!appState) {
    return (
      <div className="App">
        <Choice
          onTakeSurvey={() => {
            setAppState("survey");
          }}
          onViewResults={() => {
            setAppState("results");
          }}
        />
      </div>
    );
  }

  if (appState === "results") {
    return (
      <div className="App">
        <Results />
      </div>
    );
  }

  return (
    <div className="App">
      {currentPage == 0 && (
        <Page1
          onNextPage={handleNextPage}
          onPreviousPage={handlePreviousPage}
          setData={setUserDetails}
        />
      )}
      {currentPage == 1 && (
        <Page2
          onNextPage={handleSubmit}
          onPreviousPage={handlePreviousPage}
          setData={setMatrixData}
        />
      )}
      {currentPage == 20 && (
        <Page2a
          onNextPage={handleNextPage}
          onPreviousPage={handlePreviousPage}
        />
      )}
      {currentPage == 2 && (
        <Page3 onPreviousPage={handlePreviousPage} onSubmit />
      )}
    </div>
  );
}

export default App;
