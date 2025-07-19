import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

interface Props {
  runway: string;
}

export function ResultsCard({ runway }: Props) {
  const { data } = useQuery(['perf', runway], async () => {
    const res = await axios.post('/v1/perf/calc', { runway, temp: 15 });
    return res.data;
  });

  if (!data) return <p>Loading...</p>;
  return (
    <div>
      <h3>V1 {data.V1}</h3>
      <p>Fuel delta {data.FuelDelta}</p>
    </div>
  );
}
