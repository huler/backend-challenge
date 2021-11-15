function Choice({ onTakeSurvey, onViewResults }) {
  return (
    <>
      <h1>Acme Co employee satisfaction survey</h1>

      <button onClick={onTakeSurvey}>Take survey</button>
      <button onClick={onViewResults}>View results</button>
    </>
  );
}

export default Choice;
