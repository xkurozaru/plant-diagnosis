import { Alert, AlertIcon, Button, FormControl, FormLabel, Input, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useState } from 'react';

const ModelForm = ({ updateModelList }) => {
  const [modelData, setModelData] = useState({
    model_name: '',
    network_name: '',
    param_path: '',
    labels: [],
  });
  const [loading, setLoading] = useState(false);
  const [successAlert, setSuccessAlert] = useState(false);
  const [errorAlert, setErrorAlert] = useState(false);
  const token = Cookies.get('token');

  const sendRequest = async () => {
    setLoading(true);
    setSuccessAlert(false); // リクエストを送信する前にアラートを隠す
    setErrorAlert(false); // リクエストを送信する前にアラートを隠す
    const url = `${process.env.NEXT_PUBLIC_HOST}/api/v1/prediction/models`;
    try {
      const response = await axios.post(url, modelData, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.status === 200) {
        // リクエストが成功した場合の処理
        console.log('リクエストが成功しました');
        setModelData({
          model_name: '',
          network_name: '',
          param_path: '',
          labels: [],
        }); // フォームをリセット
        setSuccessAlert(true); // 成功アラートを表示
        updateModelList(); // 新しい関数を呼び出す
      } else {
        // リクエストが失敗した場合の処理
        console.error('リクエストが失敗しました');
        setErrorAlert(true); // エラーアラートを表示
      }
    } catch (error) {
      console.error('エラーが発生しました:', error);
      setErrorAlert(true); // エラーアラートを表示
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;

    // ラベルの場合、特別な処理を追加
    if (name === 'labels') {
      // カンマで区切られたラベルを配列に分割
      const labelsArray = value.split(',');
      setModelData((prevData) => ({
        ...prevData,
        labels: labelsArray, // ラベルを配列に設定
      }));
    } else {
      setModelData((prevData) => ({
        ...prevData,
        [name]: value,
      }));
    }
  };

  return (
    <VStack spacing={4} maxW="fit-content" align="start">
      {successAlert && (
        <Alert status="success">
          <AlertIcon />
          モデルを作成しました
        </Alert>
      )}
      {errorAlert && (
        <Alert status="error">
          <AlertIcon />
          モデルの作成に失敗しました
        </Alert>
      )}
      <Text fontSize="xl" fontWeight="bold">
        モデル作成
      </Text>
      <FormControl>
        <FormLabel>モデル名</FormLabel>
        <Input
          type="text"
          name="model_name"
          value={modelData.model_name}
          onChange={handleChange}
          required
        />
      </FormControl>
      <FormControl>
        <FormLabel>ネットワーク名</FormLabel>
        <Input
          type="text"
          name="network_name"
          value={modelData.network_name}
          onChange={handleChange}
          required
        />
      </FormControl>
      <FormControl>
        <FormLabel>パラメーターパス</FormLabel>
        <Input
          type="text"
          name="param_path"
          value={modelData.param_path}
          onChange={handleChange}
          required
        />
      </FormControl>
      <FormControl>
        <FormLabel>ラベル(カンマで区切って入力)</FormLabel>
        <Input
          type="text"
          name="labels"
          value={modelData.labels.join(',')} // 配列をカンマ区切りの文字列に変換
          onChange={handleChange}
          required
        />
      </FormControl>
      <Button colorScheme='teal' variant='solid' size="lg" onClick={sendRequest} isLoading={loading} loadingText="モデル作成中">モデル作成</Button>
    </VStack>
  );
};

export default ModelForm;
