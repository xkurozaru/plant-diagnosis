import { Alert, AlertIcon, Button, Card, CardBody, CardFooter, CardHeader, Heading, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useEffect, useState } from 'react';



const ModelList = ({ models, setModels }) => {
  const [loading, setLoading] = useState(true);
  const [successAlert, setSuccessAlert] = useState(false);
  const [errorAlert, setErrorAlert] = useState(false);
  const token = Cookies.get('token');

  useEffect(() => {
    const fetchModels = async () => {
      try {
        const url = `${process.env.NEXT_PUBLIC_HOST}/api/v1/prediction/models`;
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
    setSuccessAlert(false); // リクエストを送信する前にアラートを隠す
    setErrorAlert(false); // リクエストを送信する前にアラートを隠す

    try {
      const url = `${process.env.NEXT_PUBLIC_HOST}/api/v1/prediction/models/${modelId}`;
      const response = await axios.delete(url, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status === 200) {
        // 削除成功
        console.log('モデルを削除しました');
        const updatedModels = models.filter((model) => model.id !== modelId);
        setModels(updatedModels);
        setSuccessAlert(true);
      } else {
        console.error('モデルの削除に失敗しました');
        setErrorAlert(true);
      }
    } catch (error) {
      console.error('エラーが発生しました:', error);
    }
  };

  return (
    <VStack spacing={4} maxW="fit-content" align="start">
      {successAlert && (
        <Alert status="success">
          <AlertIcon />
          モデルを削除しました
        </Alert>
      )}
      {errorAlert && (
        <Alert status="error">
          <AlertIcon />
          モデルの削除に失敗しました
        </Alert>
      )}
      <Text fontSize="xl" fontWeight="bold">
        モデル一覧
      </Text>
      {loading ? (
        <Text>データを読み込んでいます...</Text>
      ) : (
        models.map((model) => (
          <Card key={model.id} borderRadius="md" maxW="lg">
            <CardHeader>
              <Heading size="md">{model.model_name}</Heading>
            </CardHeader>
            <CardBody>
              <Text style={{ whiteSpace: 'pre-wrap' }}>ラベル: {model.labels.join(', ')}</Text>
            </CardBody>
            <CardFooter>
              <Button colorScheme="red" onClick={() => deleteModel(model.id)}>モデル削除</Button>
            </CardFooter>
          </Card>
        ))
      )}
    </VStack>
  );
};

export default ModelList;
