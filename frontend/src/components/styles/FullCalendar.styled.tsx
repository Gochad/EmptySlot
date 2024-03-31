import styled from "styled-components";
import Modal from "react-modal";

export const Form = styled.form`
  padding: 20px;
  max-width: 400px;
  width: 100%;
  background-color: white;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;

  label {
    display: block;
    margin-bottom: 5px;
    color: #666;
  }

  input {
    width: 100%;
    padding: 10px 0;
    margin-bottom: 20px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }

  p.error {
    color: red;
    text-align: center;
  }
`;


export const ModalStyled = styled(Modal)<object>`
  min-height: 10vh;
  min-width: 400px;
  max-height: 90vh;
  max-width: 1000px;
  margin: auto;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
  border-radius: 5px;
  border: 2px solid #000;
`;


export const CalendarStyles = {
    height: '80vh',
    width: `90vw`,
}