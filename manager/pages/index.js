import UploadForm from "@/components/UploadForm";
import { Box } from "@chakra-ui/react";
import Header from "../components/Header";

export default function Home() {
  return (
    <Box>
      <Header />
      <UploadForm />
    </Box>
  );
};
