interface Props {
  value: string;
  onChange: (v: string) => void;
}

export function RunwaySelect({ value, onChange }: Props) {
  const options = ['16L', '34R'];
  return (
    <select value={value} onChange={(e) => onChange(e.target.value)}>
      {options.map((o) => (
        <option key={o}>{o}</option>
      ))}
    </select>
  );
}
