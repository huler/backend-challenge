function Choice({ onTakeSurvey, onViewResults }) {
  return (
    <>
      <h1>Acme Co satisfaction survey</h1>

      <button onClick={onTakeSurvey}>Take survey</button>
      <button onClick={onViewResults}>View results</button>
    </>
  );
}

export default Choice;
