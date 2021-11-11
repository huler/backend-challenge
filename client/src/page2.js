import React, { useState } from "react";

function Page2({ onNextPage, onPreviousPage, setData }) {
  const questions = [
    "I feel encouraged to come up with new and better ways of doing things.",
    "My work gives me a feeling of personal accomplishment.",
    "I have the tools and resources to do my job well.",
    "In my job, I have clearly defined quality goals",
  ];

  let [answer1, setAnswer1] = useState("2");
  let [answer2, setAnswer2] = useState("2");
  let [answer3, setAnswer3] = useState("2");
  let [answer4, setAnswer4] = useState("2");

  const makeRow = (question, answer, onChange) => {
    return (
      <tr
        onChange={(e) => {
          onChange(e.target.value);
        }}
      >
        <td>{question}</td>
        <td>
          <input
            type="radio"
            value={0}
            checked={answer === "0"}
            name={question}
          />
        </td>
        <td>
          <input
            type="radio"
            value={1}
            checked={answer === "1"}
            name={question}
          />
        </td>
        <td>
          <input
            type="radio"
            value={2}
            checked={answer === "2"}
            name={question}
          />
        </td>
        <td>
          <input
            type="radio"
            value={3}
            checked={answer === "3"}
            name={question}
          />
        </td>
        <td>
          <input
            type="radio"
            value={4}
            checked={answer === "4"}
            name={question}
          />
        </td>
      </tr>
    );
  };

  const handleNextPage = () => {
    setData({
      answer1,
      answer2,
      answer3,
      answer4,
    });
    onNextPage();
  };

  return (
    <>
      <h1>Acme Co satisfaction survey</h1>

      <table>
        <thead>
          <tr>
            <th></th>
            <th>Strongly disagree</th>
            <th>Somewhat disagree</th>
            <th>Neither agree nor disagree</th>
            <th>Somewhat agree</th>
            <th>Strongly agree</th>
          </tr>
        </thead>
        <tbody>
          {makeRow(questions[0], answer1, setAnswer1)}
          {makeRow(questions[1], answer2, setAnswer2)}
          {makeRow(questions[2], answer2, setAnswer3)}
          {makeRow(questions[3], answer4, setAnswer4)}
        </tbody>
      </table>

      <button onClick={onPreviousPage}>Previous page</button>
      <button onClick={handleNextPage}>Next page</button>
    </>
  );
}

export default Page2;
