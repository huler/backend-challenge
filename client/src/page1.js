import React, { useState } from "react";

function Page1({ onNextPage, onPreviousPage, setData }) {
  const departments = [
    "SAAS development",
    "Bespoke development",
    "Creative",
    "HR",
    "QA",
    "Project Management",
  ];

  let [email, setEmail] = useState("");
  const onHandleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  let [department, setDepartment] = useState(departments[0]);
  const onHandleDepartmentChange = (e) => {
    setDepartment(e.target.value);
  };

  const handleNextPage = () => {
    setData({
      email,
      department,
    });
    onNextPage();
  };

  return (
    <>
      <h1>Acme Co employee satisfaction survey</h1>
      <div>Welcome to the Huler employee satisfaction survey.</div>

      <div className="question-container" key="email">
        Email:
        <input type="text" value={email} onChange={onHandleEmailChange}></input>
      </div>
      <div className="question-container" key="department">
        Department:
        <select
          name="departments"
          value={department}
          onChange={onHandleDepartmentChange}
        >
          {departments.map((item) => {
            return <option value={item}>{item}</option>;
          })}
        </select>
      </div>
      <button onClick={handleNextPage}>Next page</button>
    </>
  );
}

export default Page1;
