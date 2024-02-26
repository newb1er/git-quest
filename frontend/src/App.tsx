import "./App.css";
import { useFetchQuests } from "./lib/hooks/api";

function QuestList() {
  const { quests, loading } = useFetchQuests("http://localhost:8080");

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <h1>Quests</h1>
      <ul>
        {quests.map((quest) => (
          <li key={quest.id}>{quest.title}</li>
        ))}
      </ul>
    </>
  );
}

function App() {
  // const { sendControlMessage } = useWebsocket("ws://localhost:8080/ws");

  return <QuestList />;
}

export default App;
