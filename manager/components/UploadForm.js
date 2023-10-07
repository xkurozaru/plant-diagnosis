// components/UploadForm.js
import {
  Alert,
  AlertIcon,
  AlertTitle,
  Box,
  Button,
  Image,
  Input,
  Spinner,
  Stack,
} from '@chakra-ui/react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useState } from 'react';

const placeholderImage = 'https://via.placeholder.com/512'; // プレースホルダー画像のURL

export default function UploadForm() {
  const [token, setToken] = useState(Cookies.get("token"));
  const [file, setFile] = useState(null);
  const [id, setId] = useState('');
  const [responseMessage, setResponseMessage] = useState(null);
  const [imagePreview, setImagePreview] = useState(null);
  const [loading, setLoading] = useState(false); // ローディング状態を管理

  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
    setFile(selectedFile);

    const reader = new FileReader();
    reader.onload = (e) => {
      setImagePreview(e.target.result);
    };
    reader.readAsDataURL(selectedFile);
  };

  const handleIdChange = (e) => {
    setId(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append('image', file);
    formData.append('id', id);

    try {
      setLoading(true); // リクエストが始まったらローディングを表示
      const response = await axios.post('http://localhost:8000/api/v1/predict', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      const responseData = response.data;
      setResponseMessage(responseData.message);
    } catch (error) {
      console.error(error);
      setResponseMessage('エラーが発生しました');
    } finally {
      setLoading(false); // リクエストが完了したらローディングを非表示
    }
  };

  return (
    <Box maxW="fit-content">
      <form onSubmit={handleSubmit}>
        {responseMessage && (
        <Alert status={responseMessage === 'Success' ? 'success' : 'error'}>
          <AlertIcon />
          <AlertTitle>{responseMessage}</AlertTitle>
        </Alert>
        )}
        <Box paddingLeft={4} marginTop={4}>
          {imagePreview ? (
            <Image src={imagePreview} alt="プレビュー" boxSize="224" borderRadius="10%" />
          ) : (
            <Image src={placeholderImage} alt="プレースホルダー" boxSize="224" borderRadius="10%" />
              )}
          <Box>
            <Input type="file" onChange={handleFileChange} />
          </Box>
          <Box>
            <Input
              type="text"
              placeholder="ID"
              value={id}
              onChange={handleIdChange}
            />
          </Box>
          <Stack spacing={10} direction="row" align="center">
            <Button type="submit">診断開始</Button>
            {loading && <Spinner size="xl" />}
          </Stack>
        </Box>
      </form>
    </Box>
  );
}
