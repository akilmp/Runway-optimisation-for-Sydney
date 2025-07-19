import { useState } from 'react';
import { SeatMap } from './components/SeatMap';
import { RunwaySelect } from './components/RunwaySelect';
import { ResultsCard } from './components/ResultsCard';

export default function App() {
  const [runway, setRunway] = useState('16L');
  return (
    <div>
      <h1>Syd Perf</h1>
      <RunwaySelect value={runway} onChange={setRunway} />
      <SeatMap />
      <ResultsCard runway={runway} />
    </div>
  );
}
