import { Button } from "@chakra-ui/button";
import { Alert, Box, FormControl, FormLabel, Image, Input, Spinner, Stack } from "@chakra-ui/react";
import axios from "axios";
import { useState } from "react";

export default function Form() {
  const [selectedImage, setSelectedImage] = useState(null);
  const [alert, setAlert] = useState(null);
  const [formData, setFormData] = useState({ image: null });
  const [loading, setLoading] = useState(null);
  const [result, setResult] = useState(null);

  // 画像をアップロードしてプレビューに表示する
  const onChangeInputFile = (e) => {
    const file = e.target.files[0]
    const reader = new FileReader()
    reader.onload = (e) => {
      console.log(e.target.result)
      setSelectedImage(e.target.result)
      setFormData({ image: file })
    }
    if (file) {
      reader.readAsDataURL(file)
    } else {
      setSelectedImage(null)
      setFormData({ image: null })
    }
  }

  // 画像を送信して診断結果を受け取る
  const handlePrediction = async () => {
    setLoading(<Spinner size="xl" />)

    if (selectedImage == null) {
      setAlert(<Alert status="error">画像を選択してください</Alert>)
      return
    }
    setAlert(null)

    try {
      const response = await axios.post(`http://localhost:8000/api/predict`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        }
      });

      setResult(<Alert status="success">診断結果: {response.data.class}</Alert>)
      console.log("サーバーからの応答:", response);
    } catch (error) {
      console.log("エラー:", error);
    } finally {
      setLoading(null)
    }
  }

  return (
    <Box as="form" px="6">
      {alert}
      {result}
      <FormControl>
        <Box>
          <FormLabel>画像を選択</FormLabel>
          <Input type="file" name="image" accept="image/*" onChange={onChangeInputFile} />
        </Box>
        <Box>
          <FormLabel>プレビュー</FormLabel>
          <Image src={selectedImage} fallbackSrc='https://via.placeholder.com/512' boxSize="224" borderRadius="10%" />
        </Box>
        <Stack spacing={4} direction="row" align="center">
          <Button colorScheme="teal" variant="outline" my="4" onClick={handlePrediction}>診断開始</Button>
          {loading}
        </Stack>
      </FormControl>
    </Box>
  );
}
