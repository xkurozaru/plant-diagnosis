import { Box, Button, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useEffect, useState } from 'react';

const ModelList = () => {
  const [models, setModels] = useState([]);
  const [loading, setLoading] = useState(true);
  const token = Cookies.get('token');

  useEffect(() => {
    const fetchModels = async () => {
      try {
        const url = 'http://localhost:8000/api/v1/prediction/models';
        const response = await axios.get(url, {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.status === 200) {
          setModels(response.data.prediction_models);
          setLoading(false);
        } else {
          console.error('リクエストが失敗しました');
          setLoading(false);
        }
      } catch (error) {
        console.error('エラーが発生しました:', error);
        setLoading(false);
      }
    };

    fetchModels();
  }, []);

  const deleteModel = async (modelId) => {
    try {
      const url = `http://localhost:8000/api/v1/prediction/models/${modelId}`;
      const response = await axios.delete(url, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status === 204) {
        // 削除成功
        const updatedModels = models.filter((model) => model.id !== modelId);
        setModels(updatedModels);
      } else {
        console.error('モデルの削除に失敗しました');
      }
    } catch (error) {
      console.error('エラーが発生しました:', error);
    }
  };

  return (
    <VStack spacing={4}>
      <Text fontSize="xl" fontWeight="bold">
        モデル一覧
      </Text>
      {loading ? (
        <Text>データを読み込んでいます...</Text>
      ) : (
        models.map((model) => (
          <Box key={model.id} borderWidth="1px" p={4} borderRadius="md">
            <Text fontSize="lg">{model.model_name}</Text>
            <Text>ラベル: {model.labels.join(', ')}</Text>
            <Button onClick={() => deleteModel(model.id)}>モデルを削除</Button>
          </Box>
        ))
      )}
    </VStack>
  );
};

export default ModelList;
