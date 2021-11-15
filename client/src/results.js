import React, { useEffect, useState } from "react";

function Results({ onNextPage, onPreviousPage }) {
  const questions = [
    "I feel encouraged to come up with new and better ways of doing things.",
    "My work gives me a feeling of personal accomplishment.",
    "I have the tools and resources to do my job well.",
    "In my job, I have clearly defined quality goals",
  ];

  const [results, setResults] = useState(null);

  const translateItem = (item) => {
    const rounded = Math.round(item);
    if (rounded === 0) return "Strongly disagree";
    if (rounded === 1) return "Somewhat disagree";
    if (rounded === 2) return "Neither agree nor disagree";
    if (rounded === 3) return "Somewhat agree";
    if (rounded === 4) return "Strongly agree";

    return "???";
  };
  const makeRow = (question, contents, heading) => {
    return (
      <tr key={question}>
        <td>{question}</td>
        {contents.map((item) => (
          <td>{heading ? item : translateItem(item)}</td>
        ))}
      </tr>
    );
  };

  console.log(results);

  useEffect(() => {
    fetch(
      "https://9qa7ws9f9l.execute-api.eu-west-1.amazonaws.com/Prod/getresults"
    )
      .then((response) => response.json())
      .then((data) => {
        setResults(data);
      });
  }, []);

  return (
    <>
      <h1>Acme Co satisfaction survey</h1>

      <h2>Average results</h2>
      {results && (
        <table>
          <thead>{makeRow("Department", Object.keys(results), true)}</thead>
          <tbody>
            {makeRow(
              questions[0],
              Object.values(results).map((item) => item[0]),
              false
            )}
            {makeRow(
              questions[1],
              Object.values(results).map((item) => item[1]),
              false
            )}
            {makeRow(
              questions[2],
              Object.values(results).map((item) => item[2]),
              false
            )}
            {makeRow(
              questions[3],
              Object.values(results).map((item) => item[3]),
              false
            )}
          </tbody>
        </table>
      )}
    </>
  );
}

export default Results;
