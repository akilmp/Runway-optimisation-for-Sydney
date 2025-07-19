import { useAppSelector } from '../hooks';

export function SeatMap() {
  const pax = useAppSelector((s) => s.seats.passengers);
  return (
    <div>
      <h2>Seat Map</h2>
      <ul>
        {pax.map((p, i) => (
          <li key={i}>{p}</li>
        ))}
      </ul>
    </div>
  );
}
