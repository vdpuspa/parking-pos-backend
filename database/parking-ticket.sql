create database parking_pos;
create table ParkingTicket{
    ParkingTicketID bigserial,
    VehicleNumber varchar(10),
    VehicleType varchar(20),
    PaymentMethod varhcar(50),
    CheckInTime timestamp without timezone,
    CheckOutTime timestamp withput timezone,
    Status varchar(3)
};
