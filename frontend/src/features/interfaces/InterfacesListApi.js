export async function fetchInterfaces() {
    const res = await fetch("http://localhost:8080/api/interfaces");
    if (!res.ok) throw new Error("Erreur API");
    return await res.json();
  }
