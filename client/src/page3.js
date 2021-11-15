import { useEffect } from "react";

function Page3({ userDetails, matrixData }) {
  useEffect(() => {
    const payload = {
      email: userDetails.email,
      department: userDetails.department,
      results: [
        parseInt(matrixData.answer1),
        parseInt(matrixData.answer2),
        parseInt(matrixData.answer3),
        parseInt(matrixData.answer4),
      ],
    };
    console.log(payload);

    fetch(
      "https://9qa7ws9f9l.execute-api.eu-west-1.amazonaws.com/Prod/postresults",
      {
        method: "POST",
        body: JSON.stringify(payload),
      }
    );
  }, []);

  return (
    <>
      <h1>Acme Co employee satisfaction survey</h1>

      <h2>Submit survey</h2>

      <div>
        Thank you for completing the survey, you can now close this page.
      </div>
    </>
  );
}

export default Page3;
