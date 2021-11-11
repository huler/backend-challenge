function Page2a({ onSubmit, onPreviousPage }) {
  const questions = [];

  return (
    <>
      <h1>Acme Co satisfaction survey</h1>

      <h2>Submit survey</h2>

      <div>Thank you for completing the survey.</div>

      <button onClick={onPreviousPage}>Previous page</button>
      <button onClick={onSubmit}>Submit</button>
    </>
  );
}

export default Page2a;
