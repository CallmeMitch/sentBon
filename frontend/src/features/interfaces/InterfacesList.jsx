import { useEffect, useState } from 'react';
import { fetchInterfaces } from './InterfacesListApi';

function InterfacesList() {
  const [interfaces, setInterfaces] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchInterfaces()
      .then(setInterfaces)
      .catch(setError);
  }, []);

  if (error) return <p>Erreur: {error.message}</p>;

  return (
    <div>
      <h2>Interfaces réseau</h2>
      <ul>
        {interfaces.map((iface, index) => (
          <li key={index}>{iface.Name}</li>
        ))}
      </ul>
    </div>
  );
}

export default InterfacesList;
