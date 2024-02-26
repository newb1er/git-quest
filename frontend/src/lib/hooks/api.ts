import { useEffect, useState } from "react";

const useApi = (baseUrl: string) => {
  const getQuests = async () => {
    const response = await fetch(baseUrl + "/quests");
    return response.json();
  };

  return {
    getQuests,
  };
};

const useFetchQuests = (baseUrl: string) => {
  const [quests, setQuests] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(baseUrl + "/quests");
      const data = await response.json();
      setQuests(data.quests);
      setLoading(false);
    };

    fetchData();
  }, [baseUrl]);

  return { quests, loading };
};

export { useApi, useFetchQuests };
